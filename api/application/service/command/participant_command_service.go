package command

import (
	"api/application/dto"
	"api/domain/entity"
	"api/domain/repository"

	"github.com/google/uuid"
)

type ParticipantCommandService struct {
	participantRepo repository.ParticipantRepositoryInterface
}

func NewParticipantCommandService(pr repository.ParticipantRepositoryInterface) *ParticipantCommandService {
	return &ParticipantCommandService{participantRepo: pr}
}

func (s *ParticipantCommandService) CreateParticipant(participant dto.ParticipantCreateRequest, userID string) (string, error) {
	var participantEntity *entity.Participant
	if participant.Role == "ai_coach" {
		participantEntity = dto.ParticipantCreateRequestToEntity(participant, uuid.New().String())
	} else {
		participantEntity = dto.ParticipantCreateRequestToEntity(participant, userID)
	}

	userID, err := s.participantRepo.Create(participantEntity)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *ParticipantCommandService) UpdateParticipant(participantID string, participant dto.ParticipantUpdateRequest) error {
	participantEntity, err := s.participantRepo.FindByID(participantID)
	if err != nil {
		return err
	}

	dto.ParticipantUpdateRequestToEntity(participantEntity, participant)

	err = s.participantRepo.Update(participantEntity)
	if err != nil {
		return err
	}

	return nil
}
