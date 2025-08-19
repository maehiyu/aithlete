package service

import (
	"api/application/dto"
)

type MockChatCommandService struct {
	CreateChatFunc func(chat dto.ChatCreateRequest) (*dto.ChatDetailResponse, error)
	UpdateChatFunc func(chat dto.ChatUpdateRequest) (*dto.ChatDetailResponse, error)
}

func NewMockChatCommandService() *MockChatCommandService {
	return &MockChatCommandService{
		CreateChatFunc: func(chat dto.ChatCreateRequest) (*dto.ChatDetailResponse, error) { panic("not used") },
		UpdateChatFunc: func(chat dto.ChatUpdateRequest) (*dto.ChatDetailResponse, error) { panic("not used") },
	}
}

func (m *MockChatCommandService) CreateChat(chat dto.ChatCreateRequest) (*dto.ChatDetailResponse, error) {
	return m.CreateChatFunc(chat)
}

func (m *MockChatCommandService) UpdateChat(chat dto.ChatUpdateRequest) (*dto.ChatDetailResponse, error) {
	return m.UpdateChatFunc(chat)
}