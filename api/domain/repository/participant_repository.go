//go:generate mockgen -source=participant_repository.go -destination=mocks/mock_participant_repository.go -package=mocks ParticipantRepositoryInterface
package repository

import (
	"api/domain/entity"
)

type ParticipantRepositoryInterface interface {
	FindByID(participantID string) (*entity.Participant, error)
	FindByIDs(participantIDs []string) ([]*entity.Participant, error)
	Create(participant *entity.Participant) (string, error)
	Update(participant *entity.Participant) error
}