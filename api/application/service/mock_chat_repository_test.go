package service

import "api/domain/entity"

type MockChatRepository struct {
	CreateChatFunc   func(request *entity.Chat) (*entity.Chat, error)
	FindChatByIDFunc func(chatId string) (*entity.Chat, error)
	UpdateChatFunc   func(chat *entity.Chat) (*entity.Chat, error)
}

func NewMockChatRepository() *MockChatRepository {
	return &MockChatRepository{
		CreateChatFunc:   func(request *entity.Chat) (*entity.Chat, error) { panic("not used") },
		FindChatByIDFunc: func(chatId string) (*entity.Chat, error) { panic("not used") },
		UpdateChatFunc:   func(chat *entity.Chat) (*entity.Chat, error) { panic("not used") },
	}
}

func (m *MockChatRepository) CreateChat(request *entity.Chat) (*entity.Chat, error) {
	return m.CreateChatFunc(request)
}
func (m *MockChatRepository) FindChatByID(chatId string) (*entity.Chat, error) {
	return m.FindChatByIDFunc(chatId)
}
func (m *MockChatRepository) UpdateChat(chat *entity.Chat) (*entity.Chat, error) {
	return m.UpdateChatFunc(chat)
}
