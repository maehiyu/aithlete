package query

import (
	"api/application/dto"
)

type MockParticipantQuery struct {
	FindParticipantsByChatIDFunc func(chatID string) ([]dto.ParticipantResponse, error)
	FindParticipantByIDFunc       func(participantID string) (*dto.ParticipantResponse, error)
}

func NewMockParticipantQuery() *MockParticipantQuery {
	return &MockParticipantQuery{
		FindParticipantsByChatIDFunc: func(chatID string) ([]dto.ParticipantResponse, error) { panic("not used") },
		FindParticipantByIDFunc:       func(participantID string) (*dto.ParticipantResponse, error) { panic("not used") },
	}
}

func (m *MockParticipantQuery) FindParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error) {
	if m.FindParticipantsByChatIDFunc != nil {
		return m.FindParticipantsByChatIDFunc(chatID)
	}
	return nil, nil
}

func (m *MockParticipantQuery) FindParticipantByID(participantID string) (*dto.ParticipantResponse, error) {
	if m.FindParticipantByIDFunc != nil {
		return m.FindParticipantByIDFunc(participantID)
	}
	return nil, nil
}