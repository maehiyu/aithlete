package query

import (
	"api/application/dto"
	"testing"
	"time"
)

func TestFindChatsByUserID(t *testing.T) {
	mockQuery := NewMockChatQuery()
	mockQuery.FindChatsByUserIDFunc = func(id string) ([]dto.ChatSummaryResponse, error) {
		title1 := "Chat 1"
		title2 := "Chat 2"
		latestQA1 := "Q1: Hello? A1: Hi!"
		latestQA2 := "Q2: Bye? A2: See you!"
		IconURL1 := "http://example.com/alice.png"
		IconURL2 := "http://example.com/bob.png"
		opponent1 := dto.OpponentResponse{ID: "op1", Name: "Alice", Role: "user", IconURL: &IconURL1}
		opponent2 := dto.OpponentResponse{ID: "op2", Name: "Bob", Role: "user", IconURL: &IconURL2}
		return []dto.ChatSummaryResponse{
			{
				ID:           "1",
				Title:        &title1,
				LastActiveAt: time.Now(),
				LatestQA:     &latestQA1,
				Opponent:     opponent1,
			},
			{
				ID:           "2",
				Title:        &title2,
				LastActiveAt: time.Now(),
				LatestQA:     &latestQA2,
				Opponent:     opponent2,
			},
		}, nil
	}

	service := &ChatQueryService{chatQuery: mockQuery}

	chats, err := service.GetChatsByUserID("user-123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(chats) != 2 {
		t.Fatalf("expected 2 chats, got %d", len(chats))
	}
}

func TestFindChatByID(t *testing.T) {
	mockQuery := NewMockChatQuery()
	mockQuery.FindChatByIDFunc = func(id string) (*dto.ChatDetailResponse, error) {
		title := "Chat 1"
		return &dto.ChatDetailResponse{
			ID:    id,
			Title: &title,
			Participants: []dto.ParticipantResponse{
				{ID: "user-123", Name: "Alice", Email: "alice@example.com", Role: "user", IconURL: nil},
			},
			Questions:    []dto.QuestionResponse{},
			Answers:      []dto.AnswerResponse{},
			StartedAt:    time.Now(),
			LastActiveAt: time.Now(),
		}, nil
	}

	service := &ChatQueryService{chatQuery: mockQuery}

	chat, err := service.GetChatByID("1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if chat.ID != "1" {
		t.Fatalf("expected chat ID to be '1', got %v", chat.ID)
	}
}
