//go:generate mockgen -source=participant_repository.go -destination=mocks/mock_participant_repository.go -package=mocks ParticipantRepositoryInterface
package repository

import (
	"api/domain/entity"
	"context"
)

type ParticipantRepositoryInterface interface {
	FindByID(ctx context.Context, participantID string) (*entity.Participant, error)
	FindByIDs(ctx context.Context, participantIDs []string) ([]*entity.Participant, error)
	Create(ctx context.Context, participant *entity.Participant) (string, error)
	Update(ctx context.Context, participant *entity.Participant) error
}