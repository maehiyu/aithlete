//go:generate mockgen -source=participant_query.go -destination=mocks/mock_participant_query.go -package=mocks ParticipantQueryInterface
package query

import (
	"api/domain/entity"
)

type ParticipantQueryInterface interface {
	FindParticipantsByIDs(ids []string) ([]entity.Participant, error)
	FindParticipantsByChatID(chatID string) ([]entity.Participant, error)
	FindParticipantByID(participantID string) (*entity.Participant, error)
	FindCoachesBySport(sport string) ([]entity.Participant, error)
}