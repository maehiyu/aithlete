package query

import (
	"api/application/dto"
	"reflect"
	"testing"
)

func TestFindParticipantsByChatID(t *testing.T) {
	mockQuery := NewMockParticipantQuery()
	aliceIcon := "http://example.com/alice.png"
	bobIcon := "http://example.com/bob.png"
	expected := []dto.ParticipantResponse{
		{ID: "1", Name: "Alice", Role: "user", Sports: []string{"サッカー"}, IconURL: &aliceIcon},
		{ID: "2", Name: "Bob", Role: "user", Sports: []string{"バスケットボール"}, IconURL: &bobIcon},
	}
	mockQuery.FindParticipantsByChatIDFunc = func(chatID string) ([]dto.ParticipantResponse, error) {
		if chatID != "chat1" {
			t.Errorf("expected chatID 'chat1', got %v", chatID)
		}
		return expected, nil
	}

	result, err := mockQuery.FindParticipantsByChatID("chat1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != len(expected) {
		t.Errorf("expected %d results, got %d", len(expected), len(result))
	}
	for i := range expected {
		if !reflect.DeepEqual(result[i], expected[i]) {
			t.Errorf("expected result[%d] = %+v, got %+v", i, expected[i], result[i])
		}
	}
}

func TestFindParticipantByID(t *testing.T) {
	mockQuery := NewMockParticipantQuery()
	expected := &dto.ParticipantResponse{
		ID:      "1",
		Name:    "Alice",
		Email:   "alice@example.com",
		Role:    "user",
		IconURL: nil,
	}
	mockQuery.FindParticipantByIDFunc = func(participantID string) (*dto.ParticipantResponse, error) {
		if participantID != "1" {
			t.Errorf("expected participantID '1', got %v", participantID)
		}
		return expected, nil
	}

	result, err := mockQuery.FindParticipantByID("1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != expected {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}
