package command

import (
	"api/application/dto"
	"api/domain/repository"

)

type ParticipantCommandService struct {
	participantRepo repository.ParticipantRepositoryInterface
}

func NewParticipantCommandService(pr repository.ParticipantRepositoryInterface) *ParticipantCommandService {
	return &ParticipantCommandService{participantRepo: pr}
}

func (s *ParticipantCommandService) CreateParticipant(participant dto.ParticipantCreateRequest, participantID string) (*dto.ParticipantResponse, error) {
	participantEntity := dto.ParticipantCreateRequestToEntity(participant, participantID)

	createdParticipant, err := s.participantRepo.Create(participantEntity)
	if err != nil {
		return nil, err
	}

	participantResp := dto.ParticipantEntityToResponse(createdParticipant)
	return &participantResp, nil
}

func (s *ParticipantCommandService) UpdateParticipant(participantID string, participant dto.ParticipantUpdateRequest) (*dto.ParticipantResponse, error) {
	participantEntity, err := s.participantRepo.FindByID(participantID)
	if err != nil {
		return nil, err
	}

	dto.ParticipantUpdateRequestToEntity(participantEntity, participant)

	updatedParticipant, err := s.participantRepo.Update(participantEntity)
	if err != nil {
		return nil, err
	}

	participantResp := dto.ParticipantEntityToResponse(updatedParticipant)
	return &participantResp, nil
}
