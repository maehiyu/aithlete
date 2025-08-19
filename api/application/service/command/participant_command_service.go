package service

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

func (s *ParticipantCommandService) FindParticipantByID(participantID string) (*dto.ParticipantResponse, error) {
	participant, err := s.participantRepo.FindByID(participantID)
	if err != nil {
		return nil, err
	}
	return ParticipantEntityToResponse(participant), nil
}

func (s *ParticipantCommandService) CreateParticipant(participant dto.ParticipantCreateRequest) (*dto.ParticipantResponse, error) {
	participantEntity := NewParticipantFromRequest(participant, uuid.NewString())

	createdParticipant, err := s.participantRepo.Create(participantEntity)
	if err != nil {
		return nil, err
	}

	return ParticipantEntityToResponse(createdParticipant), nil
}

func (s *ParticipantCommandService) UpdateParticipant(participant dto.ParticipantUpdateRequest) (*dto.ParticipantResponse, error) {
	participantEntity, err := s.participantRepo.FindByID(participant.ID)
	if err != nil {
		return nil, err
	}

	ApplyParticipantUpdateRequestToEntity(participantEntity, participant)

	updatedParticipant, err := s.participantRepo.Update(participantEntity)
	if err != nil {
		return nil, err
	}

	return ParticipantEntityToResponse(updatedParticipant), nil
}

func ParticipantEntityToResponse(p *entity.Participant) *dto.ParticipantResponse {
	return &dto.ParticipantResponse{
		ID:      p.ID,
		Name:    p.Name,
		Role:    p.Role,
		IconURL: p.IconURL,
	}
}

func NewParticipantFromRequest(dto dto.ParticipantCreateRequest, id string) *entity.Participant {
	return &entity.Participant{
		ID:      id,
		Name:    dto.Name,
		Role:    dto.Role,
		Sports:  dto.Sports,
		IconURL: dto.IconURL,
	}
}

func ApplyParticipantUpdateRequestToEntity(participant *entity.Participant, dto dto.ParticipantUpdateRequest) {
	if dto.Name != nil {
		participant.Name = *dto.Name
	}
	if dto.Role != nil {
		participant.Role = *dto.Role
	}
	if dto.Sports != nil {
		participant.Sports = dto.Sports
	}
	if dto.IconURL != nil {
		participant.IconURL = dto.IconURL
	}
}
