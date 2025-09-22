
package main

import (
	"api/application/dto"
	"api/application/service/command"
	appquery "api/application/service/query"
	infraquery "api/infrastructure/query"
	"api/infrastructure/repository"
	"api/presentation/handler"
	"api/presentation/middleware"
	"context"
	"log"
	"os"
	"strconv"
	"bytes"
	"net/http"
	"api/infrastructure/broker"
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
	       "vectorIndexConfig": {"vectorDimension": 384},
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
	createWeaviateSchema()
	redisHost := os.Getenv("BROKER_HOST")
	if redisHost == "" {
		redisHost = "broker"
	}
	redisPort := os.Getenv("BROKER_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisPassword := os.Getenv("BROKER_PASSWORD")
	redisDB := 0
	if dbStr := os.Getenv("BROKER_DB"); dbStr != "" {
		if n, err := strconv.Atoi(dbStr); err == nil {
			redisDB = n
		}
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})

	eventPublisher := broker.NewRedisChatEventBroker(redisClient, "chat_events")
	ragRequestBroker := broker.NewRedisChatEventBroker(redisClient, "rag_requests")
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	dsn := os.Getenv("PGX_DSN")
	if dsn == "" {
		dsn = "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME") + "?sslmode=disable"
	}

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	r.Use(middleware.AuthMiddleware())

	chatRepository := repository.NewChatRepository(pool)
	participantRepository := repository.NewParticipantRepository(pool)

	chatQuery := infraquery.NewChatQuery(pool)
	participantQuery := infraquery.NewParticipantQuery(pool)

	chatQueryService := appquery.NewChatQueryService(chatQuery)
	participantQueryService := appquery.NewParticipantQueryService(participantQuery)

	vectorStoreRepo := repository.NewVectorStoreRepository(
		"http://weaviate:8080",        // Weaviate endpoint
		"QAPair",                      // Class名
		"http://embedding:8001/embed", // Embedding API endpoint
	)

	chatCommandService := command.NewChatCommandService(chatRepository, participantRepository, eventPublisher, ragRequestBroker, vectorStoreRepo)

	if err := eventPublisher.SubscribeChatEvent(context.Background(), func(event dto.ChatEvent) error {
		return handler.ChatEventHandler(event, chatCommandService)
	}); err != nil {
		log.Printf("failed to subscribe chat_events: %v", err)
	}
	participantCommandService := command.NewParticipantCommandService(participantRepository)

	r.POST("/chats/:id/questions", handler.HandleSendQuestion(chatCommandService))
	r.POST("/chats/:id/answers", handler.HandleSendAnswer(chatCommandService))
	r.GET("/chats/:id", handler.HandleGetChat(chatQueryService))
	r.PUT("/chats/:id", handler.HandleUpdateChat(chatCommandService))
	r.GET("/chats", handler.HandleGetChats(chatQueryService))
	r.POST("/chats", handler.HandleCreateChat(chatCommandService))

	r.GET("/participants/me", handler.HandleGetCurrentUser(participantQueryService))
	r.GET("/participants/:id", handler.HandleGetParticipant(participantQueryService))
	r.POST("/participants", handler.HandleCreateParticipant(participantCommandService))
	r.PUT("/participants/:id", handler.HandleUpdateParticipant(participantCommandService))

	r.GET("/coaches", handler.HandleGetCoachesBySport(participantQueryService))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Go API")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
