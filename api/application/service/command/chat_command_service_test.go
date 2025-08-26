package command

import (
	"api/application/dto"
	"api/domain/entity"

	"testing"
)

func TestCreateChat_Success(t *testing.T) {
	chatRepo := NewMockChatRepository()
	chatRepo.CreateChatFunc = func(request *entity.Chat) (*entity.Chat, error) {
		return request, nil
	}

	participantSvc := NewMockParticipantRepository()
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:                  chatRepo,
		ParticipantRepo:           participantSvc,
	}

	reqTitle := "test chat"
	req := dto.ChatCreateRequest{
		Title:          &reqTitle,
		ParticipantIDs: []string{"1"},
	}

	resp, err := service.CreateChat(req)
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
	chatRepo := NewMockChatRepository()
	chatRepo.AddQuestionFunc = func(chatId string, q *entity.Question) (*entity.Chat, error) {
		return &entity.Chat{
			ID: chatId,
			Questions: []entity.Question{*q},
			ParticipantIDs: []string{"1"},
		}, nil
	}

	participantSvc := NewMockParticipantRepository()
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
	}

	req := dto.QuestionCreateRequest{
		ParticipantID: "1",
		Content:       "test question",
	}

	resp, err := service.SendQuestion("1", req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(resp.Questions) != 1 {
		t.Errorf("expected 1 question, got %d", len(resp.Questions))
	}
	if resp.Questions[0].Content != "test question" {
		t.Errorf("expected question content 'test question', got %v", resp.Questions[0].Content)
	}
}

func TestSendAnswer_Success(t *testing.T) {
	chatRepo := NewMockChatRepository()
	chatRepo.AddAnswerFunc = func(chatId string, a *entity.Answer) (*entity.Chat, error) {
		return &entity.Chat{
			ID: chatId,
			Answers: []entity.Answer{*a},
			ParticipantIDs: []string{"1"},
		}, nil
	}

	participantSvc := NewMockParticipantRepository()
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:        chatRepo,
		ParticipantRepo: participantSvc,
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
	if len(resp.Answers) != 1 {
		t.Errorf("expected 1 answer, got %d", len(resp.Answers))
	}
	if resp.Answers[0].Content != "test answer" {
		t.Errorf("expected answer content 'test answer', got %v", resp.Answers[0].Content)
	}
}

func TestUpdateChat_Success(t *testing.T) {
	chatRepo := NewMockChatRepository()
	chatRepo.FindChatByIDFunc = func(chatId string) (*entity.Chat, error) {
		oldTitle := "old title"
		return &entity.Chat{ID: chatId, Title: &oldTitle}, nil
	}
	chatRepo.UpdateChatFunc = func(chat *entity.Chat) (*entity.Chat, error) {
		return chat, nil
	}

	participantSvc := NewMockParticipantRepository()
	participantSvc.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return &entity.Participant{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:                  chatRepo,
		ParticipantRepo:           participantSvc,
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