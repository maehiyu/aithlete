package query

import (
	"api/application/dto"
	"api/domain/entity" // entityパッケージをインポート
	"testing"
	"time"
	"context"

	"api/application/query/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetChatsByUserID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockChatQuery := mocks.NewMockChatQueryInterface(ctrl)

	title1 := "Chat 1"
	title2 := "Chat 2"
	latestQA1 := "Q1: Hello? A1: Hi!"
	latestQA2 := "Q2: Bye? A2: See you!"
	IconURL1 := "http://example.com/alice.png"
	IconURL2 := "http://example.com/bob.png"
	opponent1 := dto.OpponentResponse{ID: "op1", Name: "Alice", Role: "user", IconURL: &IconURL1}
	opponent2 := dto.OpponentResponse{ID: "op2", Name: "Bob", Role: "user", IconURL: &IconURL2}

	expectedChats := []dto.ChatSummaryResponse{
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
	}

	mockChatQuery.EXPECT().FindChatsByUserID(gomock.Any(), gomock.Any()).Return(expectedChats, nil).Times(1)

	service := NewChatQueryService(mockChatQuery, nil) // ParticipantQueryInterfaceは今回は不要なのでnil

	chats, err := service.GetChatsByUserID(ctx, "user-123")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(chats) != 2 {
		t.Fatalf("expected 2 chats, got %d", len(chats))
	}

}

func TestGetChatByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()

	mockChatQuery := mocks.NewMockChatQueryInterface(ctrl)
	mockParticipantQuery := mocks.NewMockParticipantQueryInterface(ctrl)

	// chatQuery.FindChatByID のモック設定
	chatTitle := "Test Chat"
	mockChatEntity := &entity.Chat{
		ID:             "chat-1",
		Title:          &chatTitle,
		ParticipantIDs: []string{"user-1", "user-2"},
		StartedAt:      time.Now(),
		LastActiveAt:   time.Now(),
	}
	mockChatQuery.EXPECT().FindChatByID(gomock.Any(),gomock.Eq("chat-1")).Return(mockChatEntity, nil).Times(1)

	// participantQuery.FindParticipantsByIDs のモック設定
	mockParticipantEntities := []entity.Participant{
		{ID: "user-1", Name: "Alice", Email: "alice@example.com", Role: "user"},
		{ID: "user-2", Name: "Bob", Email: "bob@example.com", Role: "user"},
	}
	mockParticipantQuery.EXPECT().FindParticipantsByIDs(gomock.Any(), gomock.Eq([]string{"user-1", "user-2"})).Return(mockParticipantEntities, nil).Times(1)

	// dto.ChatEntityToDetailResponse は直接呼び出される関数なので、ここではモック化せず実関数を呼び出す
	// ChatQueryService のテストなので、この変換関数が正しく動作することは前提とする

	service := NewChatQueryService(mockChatQuery, mockParticipantQuery)

	chatDetail, err := service.GetChatByID(ctx, "chat-1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if chatDetail == nil {
		t.Fatal("expected chat detail, got nil")
	}
	if chatDetail.ID != "chat-1" {
		t.Errorf("expected chat ID 'chat-1', got %s", chatDetail.ID)
	}
	if *chatDetail.Title != *mockChatEntity.Title {
		t.Errorf("expected chat title '%s', got '%s'", *mockChatEntity.Title, *chatDetail.Title)
	}
	if len(chatDetail.Participants) != 2 {
		t.Errorf("expected 2 participants, got %d", len(chatDetail.Participants))
	}
	// 必要に応じて、chatDetail の内容を詳細に検証
}