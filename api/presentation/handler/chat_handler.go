package handler

import (
	"net/http"

	"api/application/dto"
	"api/application/service/command"
	"api/application/service/query"

	"github.com/gin-gonic/gin"
)

func HandleGetChats(chatQueryService *query.ChatQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		chats, err := chatQueryService.GetChatsByUserID(uidStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chats)
	}
}

func HandleGetChat(chatQueryService *query.ChatQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("id")
		chat, err := chatQueryService.GetChatByID(chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chat)
	}
}

func HandleCreateChat(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		chat, err := chatCommandService.CreateChat(req, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, chat)
	}
}

func HandleUpdateChat(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("id")
		var req dto.ChatUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		chat, err := chatCommandService.UpdateChat(chatID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, chat)
	}
}

func HandleSendQuestion(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("id")
		var req dto.QuestionCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		question, err := chatCommandService.SendQuestion(chatID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, question)
	}
}

func HandleSendAnswer(chatCommandService *command.ChatCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		chatID := c.Param("id")
		var req dto.AnswerCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		answer, err := chatCommandService.SendAnswer(chatID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, answer)
	}
}