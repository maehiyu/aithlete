package main

import (
	"api/application/service/command"
	appquery "api/application/service/query"
	infraquery "api/infrastructure/query"
	"api/infrastructure/repository"
	"api/presentation/handler"
	"api/presentation/middleware"
	"context"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
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

	chatCommandService := command.NewChatCommandService(chatRepository, participantRepository)
	participantCommandService := command.NewParticipantCommandService(participantRepository)

	r.POST("/chats/:id/questions", handler.HandleSendQuestion(chatCommandService))
	r.POST("/chats/:id/answers", handler.HandleSendAnswer(chatCommandService))
	r.GET("/chats/:id", handler.HandleGetChat(chatQueryService))
	r.PUT("/chats/:id", handler.HandleUpdateChat(chatCommandService))
	r.GET("/chats", handler.HandleGetChats(chatQueryService))
	r.POST("/chats", handler.HandleCreateChat(chatCommandService))
	r.GET("/participants/:id", handler.HandleGetParticipant(participantQueryService))
	r.POST("/participants", handler.HandleCreateParticipant(participantCommandService))
	r.PUT("/participants/:id", handler.HandleUpdateParticipant(participantCommandService))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Go API!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
