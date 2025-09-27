package handler

import (
	"net/http"

	"api/application/dto"
	"api/application/service/command"
	"api/application/service/query"

	"github.com/gin-gonic/gin"
)

func HandleGetCurrentUser(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		userID, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		participant, err := participantQueryService.GetParticipantByID(ctx, userID.(string))
		if err != nil {
			if err.Error() == "no rows in result set" {
				c.JSON(http.StatusOK, nil)
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participant)
	}
}

func HandleGetParticipants(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		chatID := c.Param("chat_id")
		participants, err := participantQueryService.GetParticipantsByChatID(ctx, chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participants)
	}
}

func HandleGetParticipant(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		participantID := c.Param("id")
		participant, err := participantQueryService.GetParticipantByID(ctx, participantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participant)
	}
}

func HandleCreateParticipant(participantCommandService *command.ParticipantCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		var req dto.ParticipantCreateRequest
		userID, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id, err := participantCommandService.CreateParticipant(ctx, req, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id": id})
	}
}

func HandleUpdateParticipant(participantCommandService *command.ParticipantCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		participantID := c.Param("id")
		var req dto.ParticipantUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := participantCommandService.UpdateParticipant(ctx, participantID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, nil)
	}
}

func HandleGetCoachesBySport(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		sport := c.Query("sport")
		if sport == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sport is required"})
			return
		}
		coaches, err := participantQueryService.GetCoachesBySport(ctx, sport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, coaches)
	}
}
