package handler

import (
	"net/http"
	"strings"

	"api/application/service"

	"github.com/gin-gonic/gin"
)

func HandleGetChats(chatQueryService *service.ChatQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userID := strings.TrimPrefix(auth, "Bearer ")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing JWT"})
			return
		}
		chats, err := chatQueryService.FindChatsByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chats)
	}
}
