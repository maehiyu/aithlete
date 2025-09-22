package repository

import (
	"api/domain/entity"
)

type ParticipantRepositoryInterface interface {
	FindByID(participantID string) (*entity.Participant, error)
	Create(participant *entity.Participant) (string, error)
	Update(participant *entity.Participant) error
}