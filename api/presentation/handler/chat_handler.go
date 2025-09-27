package handler

import (
	"database/sql"
	"net/http"

	"api/application/dto"
	"api/application/service/command"
	"api/application/service/query"

	"github.com/gin-gonic/gin"
)

func HandleGetChats(chatQueryService *query.ChatQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found"})
			return
		}
		uidStr, ok := userID.(string)
		if !ok || uidStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userId invalid"})
			return
		}
		chats, err := chatQueryService.GetChatsByUserID(ctx, uidStr)
		if err != nil && err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, []interface{}{})
			return
		}
		c.JSON(http.StatusOK, chats)
	}
}

func HandleGetChat(chatQueryService *query.ChatQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		chatID := c.Param("id")
		chat, err := chatQueryService.GetChatByID(ctx, chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chat)
	}
}

func HandleCreateChat(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var req dto.ChatCreateRequest
		userID, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "userId not found"})
			return
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		chatID, err := chatCommandService.CreateChat(ctx, req, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"id": chatID})
	}
}

func HandleUpdateChat(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		chatID := c.Param("id")
		var req dto.ChatUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := chatCommandService.UpdateChat(ctx, chatID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}

func HandleSendMessage(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		chatID := c.Param("id")
		var req dto.ChatItemRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := chatCommandService.SendMessage(ctx, chatID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, nil)
	}
}