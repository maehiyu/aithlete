package command

import (
	"api/application/dto"
	"api/domain/entity"
	"api/domain/repository"
	"context"

	"github.com/google/uuid"
)

type ParticipantCommandService struct {
	participantRepo repository.ParticipantRepositoryInterface
}

func NewParticipantCommandService(pr repository.ParticipantRepositoryInterface) *ParticipantCommandService {
	return &ParticipantCommandService{participantRepo: pr}
}

func (s *ParticipantCommandService) CreateParticipant(ctx context.Context, participant dto.ParticipantCreateRequest, userID string) (string, error) {
	var participantEntity *entity.Participant
	if participant.Role == "ai_coach" {
		participantEntity = dto.ParticipantCreateRequestToEntity(participant, uuid.New().String())
	} else {
		participantEntity = dto.ParticipantCreateRequestToEntity(participant, userID)
	}

	id, err := s.participantRepo.Create(ctx, participantEntity)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *ParticipantCommandService) UpdateParticipant(ctx context.Context, participantID string, participant dto.ParticipantUpdateRequest) error {
	participantEntity, err := s.participantRepo.FindByID(ctx, participantID)
	if err != nil {
		return err
	}

	dto.ParticipantUpdateRequestToEntity(participantEntity, participant)

	err = s.participantRepo.Update(ctx, participantEntity)
	if err != nil {
		return err
	}

	return nil
}
