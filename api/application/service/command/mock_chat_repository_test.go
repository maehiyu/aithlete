package command

import "api/domain/entity"

type MockChatRepository struct {
	CreateChatFunc   func(request *entity.Chat) (*entity.Chat, error)
	FindChatByIDFunc func(chatId string) (*entity.Chat, error)
	UpdateChatFunc   func(chat *entity.Chat) (*entity.Chat, error)
	AddAnswerFunc    func(chatId string, answer *entity.Answer) error
	AddQuestionFunc  func(chatId string, question *entity.Question) error
	GetParticipantIDsByChatIDFunc func(chatID string) ([]string, error)
}

func NewMockChatRepository() *MockChatRepository {
	return &MockChatRepository{
		CreateChatFunc:   func(request *entity.Chat) (*entity.Chat, error) { panic("not used") },
		FindChatByIDFunc: func(chatId string) (*entity.Chat, error) { panic("not used") },
		UpdateChatFunc:   func(chat *entity.Chat) (*entity.Chat, error) { panic("not used") },
		AddAnswerFunc:    func(chatId string, answer *entity.Answer) error { panic("not used") },
		AddQuestionFunc:  func(chatId string, question *entity.Question) error { panic("not used") },
		GetParticipantIDsByChatIDFunc: func(chatID string) ([]string, error) { return []string{"1", "2"}, nil },
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
func (m *MockChatRepository) AddAnswer(chatId string, answer *entity.Answer) error {
	return m.AddAnswerFunc(chatId, answer)
}
func (m *MockChatRepository) AddQuestion(chatId string, question *entity.Question) error {
	return m.AddQuestionFunc(chatId, question)
}
func (m *MockChatRepository) GetParticipantIDsByChatID(chatID string) ([]string, error) {
	if m.GetParticipantIDsByChatIDFunc != nil {
		return m.GetParticipantIDsByChatIDFunc(chatID)
	}
	return []string{"1", "2"}, nil
}