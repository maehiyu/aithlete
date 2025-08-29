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
		userID, exists := c.Get("userId")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		participant, err := participantQueryService.GetParticipantByID(userID.(string))
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
		chatID := c.Param("chat_id")
		participants, err := participantQueryService.GetParticipantsByChatID(chatID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participants)
	}
}

func HandleGetParticipant(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		participantID := c.Param("id")
		participant, err := participantQueryService.GetParticipantByID(participantID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participant)
	}
}

func HandleCreateUser(participantCommandService *command.ParticipantCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		participant, err := participantCommandService.CreateParticipant(req, userID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participant)
	}
}

func HandleUpdateParticipant(participantCommandService *command.ParticipantCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		participantID := c.Param("id")
		var req dto.ParticipantUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		participant, err := participantCommandService.UpdateParticipant(participantID, req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, participant)
	}
}

func HandleGetCoachesBySport(participantQueryService *query.ParticipantQueryService) gin.HandlerFunc {
	return func(c *gin.Context) {
		sport := c.Query("sport")
		if sport == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "sport is required"})
			return
		}
		coaches, err := participantQueryService.GetCoachesBySport(sport)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, coaches)
	}
}
