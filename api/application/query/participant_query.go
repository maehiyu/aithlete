//go:generate mockgen -source=participant_query.go -destination=mocks/mock_participant_query.go -package=mocks ParticipantQueryInterface
package query

import (
	"api/domain/entity"
	"context"
)

type ParticipantQueryInterface interface {
	FindParticipantsByIDs(ctx context.Context, ids []string) ([]entity.Participant, error)
	FindParticipantsByChatID(ctx context.Context, chatID string) ([]entity.Participant, error)
	FindParticipantByID(ctx context.Context, participantID string) (*entity.Participant, error)
	FindCoachesBySport(ctx context.Context, sport string) ([]entity.Participant, error)
}