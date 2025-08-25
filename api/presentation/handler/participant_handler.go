package handler

import (
	"net/http"

	"api/application/service/query"
	"api/application/service/command"
	"api/application/dto"

	"github.com/gin-gonic/gin"
)

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

func HandleCreateParticipant(participantCommandService *command.ParticipantCommandService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.ParticipantCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		participant, err := participantCommandService.CreateParticipant(req)
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
