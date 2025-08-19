package service

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

	participantSvc := NewMockParticipantCommandService()
	participantSvc.FindParticipantByIDFunc = func(id string) (*dto.ParticipantResponse, error) {
		return &dto.ParticipantResponse{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:                  chatRepo,
		ParticipantCommandService: participantSvc,
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

func TestUpdateChat_Success(t *testing.T) {
	chatRepo := NewMockChatRepository()
	chatRepo.FindChatByIDFunc = func(chatId string) (*entity.Chat, error) {
		oldTitle := "old title"
		return &entity.Chat{ID: chatId, Title: &oldTitle}, nil
	}
	
	participantSvc := NewMockParticipantCommandService()
	participantSvc.FindParticipantByIDFunc = func(id string) (*dto.ParticipantResponse, error) {
		return &dto.ParticipantResponse{
			ID:   id,
			Name: "testuser",
		}, nil
	}

	service := &ChatCommandService{
		chatRepo:                  chatRepo,
		ParticipantCommandService: participantSvc,
	}

	reqTitle := "new titile"
	req := dto.ChatUpdateRequest{
		Title:          &reqTitle,
		ParticipantIDs: []string{"1"},
	}

	resp, err := service.UpdateChat(req)
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