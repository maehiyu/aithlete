package service

import (
	"api/application/dto"
)

type MockParticipantCommandService struct {
	FindParticipantByIDFunc func(participantID string) (*dto.ParticipantResponse, error)
	CreateParticipantFunc    func(participant dto.ParticipantCreateRequest) (*dto.ParticipantResponse, error)
	UpdateParticipantFunc func(participant dto.ParticipantUpdateRequest) (*dto.ParticipantResponse, error)
}

func NewMockParticipantCommandService() *MockParticipantCommandService {
	return &MockParticipantCommandService{
		FindParticipantByIDFunc: func(participantID string) (*dto.ParticipantResponse, error) { panic("not used") },
		CreateParticipantFunc:    func(participant dto.ParticipantCreateRequest) (*dto.ParticipantResponse, error) { panic("not used") },
		UpdateParticipantFunc: func(participant dto.ParticipantUpdateRequest) (*dto.ParticipantResponse, error) { panic("not used") },
	}
}

func (m *MockParticipantCommandService) FindParticipantByID(participantID string) (*dto.ParticipantResponse, error) {
	return m.FindParticipantByIDFunc(participantID)
}

func (m *MockParticipantCommandService) CreateParticipant(participant dto.ParticipantCreateRequest) (*dto.ParticipantResponse, error) {
	return m.CreateParticipantFunc(participant)
}

func (m *MockParticipantCommandService) UpdateParticipant(participant dto.ParticipantUpdateRequest) (*dto.ParticipantResponse, error) {
	return m.UpdateParticipantFunc(participant)
}