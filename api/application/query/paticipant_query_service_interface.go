package query

import "api/application/dto"

type ParticipantQueryServiceInterface interface {
	FindParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error)
}
