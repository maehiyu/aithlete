package query

import "api/application/dto"

type ParticipantQueryInterface interface {
	FindParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error)
	FindParticipantByID(participantID string) (*dto.ParticipantResponse, error)
}
