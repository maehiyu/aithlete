package main

import (
	"api/application/service"
	"api/domain/entity"
	"api/presentation/handler"
	"api/presentation/middleware"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	// DB接続設定
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// テーブル自動生成（マイグレーション）
	err = db.AutoMigrate(&entity.Chat{}, &entity.Question{}, &entity.Answer{}, &entity.Attachment{}, &entity.PoseData{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	r.Use(middleware.AuthMiddleware())

	// ChatQueryServiceの生成
	// ここではquery層の実装を仮でnilにしています。実装に合わせて差し替えてください。
	chatQueryService := service.NewChatQueryService(nil)

	r.GET("/chats", handler.HandleGetChats(chatQueryService))

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Go API!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}
