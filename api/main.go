package main

import (
	"api/application/dto"
	"api/application/service/command"
	appquery "api/application/service/query"
	"api/infrastructure/broker"
	infraquery "api/infrastructure/query"
	"api/infrastructure/repository"
	"api/presentation/handler"
	"api/presentation/middleware"
	"bytes"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

// QAPairスキーマ作成関数
func createWeaviateSchema() {
	schema := `{
	       "class": "QAPair",
	       "vectorizer": "none",
	       "properties": [
		       {"name": "question", "dataType": ["text"]},
		       {"name": "answer", "dataType": ["text"]}
	       ]
       }`
	resp, err := http.Post("http://weaviate:8080/v1/schema", "application/json", bytes.NewBuffer([]byte(schema)))
	if err != nil {
		log.Printf("failed to create Weaviate schema: %v", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Weaviate schema creation status: %v", resp.Status)
}

func main() {
	// 1. アプリケーションのライフサイクルを管理するcontextを生成
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	createWeaviateSchema()

	// --- Redis Clientの初期化 ---
	redisHost := os.Getenv("BROKER_HOST")
	if redisHost == "" {
		redisHost = "broker"
	}
	redisPort := os.Getenv("BROKER_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: os.Getenv("BROKER_PASSWORD"),
		DB:       0, // 必要に応じて環境変数から取得
	})

	// --- データベース(PostgreSQL)接続プールの初期化 ---
	dsn := os.Getenv("PGX_DSN")
	if dsn == "" {
		dsn = "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	}
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer pool.Close()

	// --- DI (依存性注入) コンテナの構築 ---
	// Repositories & Queries
	chatRepository := repository.NewChatRepository(pool)
	participantRepository := repository.NewParticipantRepository(pool)
	appointmentRepository := repository.NewAppointmentRepository(pool)
	vectorStoreRepo := repository.NewVectorStoreRepository("http://weaviate:8080", "QAPair", "http://embedding:8001/embed")

	chatQuery := infraquery.NewChatQuery(pool)
	participantQuery := infraquery.NewParticipantQuery(pool)
	appointmentQuery := infraquery.NewAppointmentQuery(pool)

	// Brokers
	eventPublisher := broker.NewRedisChatEventBroker(redisClient, "chat_events")
	ragRequestBroker := broker.NewRedisChatEventBroker(redisClient, "rag_requests")

	// Services
	chatCommandService := command.NewChatCommandService(chatRepository, participantRepository, eventPublisher, ragRequestBroker, vectorStoreRepo)
	participantCommandService := command.NewParticipantCommandService(participantRepository)
	appointmentCommandService := command.NewAppointmentCommandService(appointmentRepository)

	chatQueryService := appquery.NewChatQueryService(chatQuery, participantQuery)
	participantQueryService := appquery.NewParticipantQueryService(participantQuery)
	appointmentQueryService := appquery.NewAppointmentQueryService(appointmentQuery)

	// --- イベントリスナーの起動 ---
	go func() {
		if err := eventPublisher.SubscribeChatEvent(ctx, func(event dto.ChatEvent) error {
			return handler.ChatEventHandler(ctx, event, chatCommandService)
		}); err != nil {
			log.Printf("failed to subscribe chat_events: %v", err)
		}
	}()

	// --- HTTPサーバー(Gin)のセットアップ ---
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	r.Use(middleware.AuthMiddleware())

	// --- ルーティング設定 ---
	r.GET("/", func(c *gin.Context) { c.String(200, "Hello, Go API") })

	// Chat routes
	r.POST("/chats/:id/messages", handler.HandleSendMessage(chatCommandService))
	r.GET("/chats/:id", handler.HandleGetChat(chatQueryService))
	r.PUT("/chats/:id", handler.HandleUpdateChat(chatCommandService))
	r.GET("/chats", handler.HandleGetChats(chatQueryService))
	r.POST("/chats", handler.HandleCreateChat(chatCommandService))

	// Participant routes
	r.GET("/participants/me", handler.HandleGetCurrentUser(participantQueryService))
	r.GET("/participants/:id", handler.HandleGetParticipant(participantQueryService))
	r.POST("/participants", handler.HandleCreateParticipant(participantCommandService))
	r.PUT("/participants/:id", handler.HandleUpdateParticipant(participantCommandService))
	r.GET("/coaches", handler.HandleGetCoachesBySport(participantQueryService))

	// Appointment routes (新規追加)
	h := handler.NewAppointmentHandler(appointmentCommandService, appointmentQueryService)
	r.POST("/appointments", h.HandleCreateAppointment())
	r.GET("/appointments/:id", h.HandleGetAppointmentByID())
	r.GET("/appointments", h.HandleListAppointments())
	r.PUT("/appointments/:id", h.HandleUpdateAppointment())
	r.DELETE("/appointments/:id", h.HandleDeleteAppointment())

	// --- HTTPサーバーの起動とグレースフルシャットダウン ---
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// アプリケーション終了シグナルを待機
	<-ctx.Done()
	stop() // コンテキストのキャンセルを即座に他に通知

	log.Println("Shutting down server...")

	// 5秒のタイムアウト付きでサーバーをグレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
