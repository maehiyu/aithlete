package query

import (
	"api/application/dto"
)

type MockChatQueryService struct {
	FindChatsByUserIDFunc func(id string) ([]dto.ChatSummaryResponse, error)
	FindChatByIDFunc      func(id string) (*dto.ChatDetailResponse, error)
}

func NewMockChatQuery() *MockChatQueryService {
	return &MockChatQueryService{
		FindChatsByUserIDFunc: func(id string) ([]dto.ChatSummaryResponse, error) { panic("not used") },
		FindChatByIDFunc:      func(id string) (*dto.ChatDetailResponse, error) { panic("not used") },
	}
}

func (m *MockChatQueryService) FindChatsByUserID(id string) ([]dto.ChatSummaryResponse, error) {
	return m.FindChatsByUserIDFunc(id)
}

func (m *MockChatQueryService) FindChatByID(id string) (*dto.ChatDetailResponse, error) {
	return m.FindChatByIDFunc(id)
}