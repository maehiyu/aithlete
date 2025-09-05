package command

import (
	"api/application/dto"
	"api/domain/entity"
	"testing"
)

// --- 共通モック定義 ---
type MockChatRepository struct {
	CreateChatFunc                func(request *entity.Chat) (*entity.Chat, error)
	GetParticipantIDsByChatIDFunc func(chatID string) ([]string, error)
	AddQuestionFunc               func(chatId string, q *entity.Question) error
	AddAnswerFunc                 func(chatId string, a *entity.Answer) error
	FindChatByIDFunc              func(chatId string) (*entity.Chat, error)
	UpdateChatFunc                func(chat *entity.Chat) (*entity.Chat, error)
	GetQuestionContentFunc        func(questionID string) (string, error)
}

func (m *MockChatRepository) GetQuestionContent(questionID string) (string, error) {
	if m.GetQuestionContentFunc != nil {
		return m.GetQuestionContentFunc(questionID)
	}
	return "", nil
}

func (m *MockChatRepository) CreateChat(request *entity.Chat) (*entity.Chat, error) {
	if m.CreateChatFunc != nil {
		return m.CreateChatFunc(request)
	}
	return request, nil
}
func (m *MockChatRepository) GetParticipantIDsByChatID(chatID string) ([]string, error) {
	if m.GetParticipantIDsByChatIDFunc != nil {
		return m.GetParticipantIDsByChatIDFunc(chatID)
	}
	return []string{}, nil
}
func (m *MockChatRepository) AddQuestion(chatId string, q *entity.Question) error {
	if m.AddQuestionFunc != nil {
		return m.AddQuestionFunc(chatId, q)
	}
	return nil
}
func (m *MockChatRepository) AddAnswer(chatId string, a *entity.Answer) error {
	if m.AddAnswerFunc != nil {
		return m.AddAnswerFunc(chatId, a)
	}
	return nil
}
func (m *MockChatRepository) FindChatByID(chatId string) (*entity.Chat, error) {
	if m.FindChatByIDFunc != nil {
		return m.FindChatByIDFunc(chatId)
	}
	return &entity.Chat{ID: chatId}, nil
}
func (m *MockChatRepository) UpdateChat(chat *entity.Chat) (*entity.Chat, error) {
	if m.UpdateChatFunc != nil {
		return m.UpdateChatFunc(chat)
	}
	return chat, nil
}

// --- テスト本体は既存のまま、NewMock...をMock...に置換 ---

func TestCreateChat_Success(t *testing.T) {
	chatRepo := &MockChatRepository{}
	chatRepo.CreateChatFunc = func(request *entity.Chat) (*entity.Chat, error) {
		return request, nil
	}
	chatRepo.GetParticipantIDsByChatIDFunc = func(chatID string) ([]string, error) {
		return []string{"1", "2"}, nil
	}

	participantSvc := &MockParticipantRepository{}
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
		EventPublisher:  nil,
	}

	reqTitle := "test chat"
	req := dto.ChatCreateRequest{
		Title:          &reqTitle,
		ParticipantIDs: []string{"1"},
	}

	resp, err := service.CreateChat(req, "1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Title != req.Title {
		t.Errorf("expected title %v, got %v", *req.Title, *resp.Title)
	}
	if len(resp.Participants) != 1 {
		t.Errorf("expected 1 participant, got %d", len(resp.Participants))
	}
	if resp.Participants[0].ID != "1" {
		t.Errorf("expected participant ID '1', got %v", resp.Participants[0].ID)
	}
}

func TestSendQuestion_Success(t *testing.T) {
	chatRepo := &MockChatRepository{}
	chatRepo.AddQuestionFunc = func(chatId string, q *entity.Question) error {
		return nil
	}
	chatRepo.GetParticipantIDsByChatIDFunc = func(chatID string) ([]string, error) {
		return []string{"1", "2"}, nil
	}

	participantSvc := &MockParticipantRepository{}
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
		EventPublisher:  nil,
	}

	req := dto.QuestionCreateRequest{
		ParticipantID: "1",
		Content:       "test question",
	}

	resp, err := service.SendQuestion("1", req, "test-token")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Content != "test question" {
		t.Errorf("expected question content 'test question', got %v", resp.Content)
	}
}

func TestSendAnswer_Success(t *testing.T) {
	chatRepo := &MockChatRepository{}
	chatRepo.AddAnswerFunc = func(chatId string, a *entity.Answer) error {
		return nil
	}
	chatRepo.GetParticipantIDsByChatIDFunc = func(chatID string) ([]string, error) {
		return []string{"1", "2"}, nil
	}

	participantSvc := &MockParticipantRepository{}
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
		EventPublisher:  nil,
	}

	req := dto.AnswerCreateRequest{
		QuestionID:    "q1",
		ParticipantID: "1",
		Content:       "test answer",
	}

	resp, err := service.SendAnswer("1", req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Content != "test answer" {
		t.Errorf("expected answer content 'test answer', got %v", resp.Content)
	}
}

func TestUpdateChat_Success(t *testing.T) {
	chatRepo := &MockChatRepository{}
	chatRepo.FindChatByIDFunc = func(chatId string) (*entity.Chat, error) {
		oldTitle := "old title"
		return &entity.Chat{ID: chatId, Title: &oldTitle}, nil
	}
	chatRepo.UpdateChatFunc = func(chat *entity.Chat) (*entity.Chat, error) {
		return chat, nil
	}
	chatRepo.GetParticipantIDsByChatIDFunc = func(chatID string) ([]string, error) {
		return []string{"1", "2"}, nil
	}

	participantSvc := &MockParticipantRepository{}
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
		EventPublisher:  nil,
	}

	reqTitle := "new titile"
	req := dto.ChatUpdateRequest{
		Title:          &reqTitle,
		ParticipantIDs: []string{"1"},
	}

	resp, err := service.UpdateChat("1", req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Title != req.Title {
		t.Errorf("expected title %v, got %v", *req.Title, *resp.Title)
	}
	if len(resp.Participants) != 1 {
		t.Errorf("expected 1 participant, got %d", len(resp.Participants))
	}
	if resp.Participants[0].ID != "1" {
		t.Errorf("expected participant ID '1', got %v", resp.Participants[0].ID)
	}
}
