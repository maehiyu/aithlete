package repository

import (
	"api/domain/entity"
)

type ParticipantRepositoryInterface interface {
	FindByID(participantID string) (*entity.Participant, error)
	Create(participant *entity.Participant) (*entity.Participant, error)
	Update(participant *entity.Participant) (*entity.Participant, error)
}