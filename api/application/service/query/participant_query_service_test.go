package query

import (
	"api/application/dto"
	"api/domain/entity"
	"context"
	"reflect"
	"testing"

	"api/application/query/mocks"
	"github.com/golang/mock/gomock"
)

func TestGetParticipantsByChatID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParticipantQuery := mocks.NewMockParticipantQueryInterface(ctrl)
	ctx := context.Background()

	// participantQuery.FindParticipantsByChatID のモック設定
	mockParticipantEntities := []entity.Participant{
		{ID: "user-1", Name: "Alice", Email: "alice@example.com", Role: "user"},
		{ID: "user-2", Name: "Bob", Email: "bob@example.com", Role: "user"},
	}
	mockParticipantQuery.EXPECT().FindParticipantsByChatID(gomock.Any(), gomock.Eq("chat-1")).Return(mockParticipantEntities, nil).Times(1)

	// dto.ParticipantEntityToResponse は直接呼び出される関数なので、ここではモック化せず実関数を呼び出す

	service := NewParticipantQueryService(mockParticipantQuery)

	participants, err := service.GetParticipantsByChatID(ctx, "chat-1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(participants) != 2 {
		t.Fatalf("expected 2 participants, got %d", len(participants))
	}

	// 期待するDTOのリストを作成
	expectedDTOs := []dto.ParticipantResponse{
		*dto.ParticipantEntityToResponse(&mockParticipantEntities[0]),
		*dto.ParticipantEntityToResponse(&mockParticipantEntities[1]),
	}

	// 詳細検証: 返されたDTOのリストが期待するリストと一致するか
	if !reflect.DeepEqual(participants, expectedDTOs) {
		t.Errorf("returned participants do not match expected participants.\nExpected: %+v\nGot: %+v", expectedDTOs, participants)
	}
}

func TestGetParticipantByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParticipantQuery := mocks.NewMockParticipantQueryInterface(ctrl)
	ctx := context.Background()

	// participantQuery.FindParticipantByID のモック設定 (見つかるケース)
	mockParticipantEntity := &entity.Participant{
		ID: "user-1", Name: "Alice", Email: "alice@example.com", Role: "user",
	}
	mockParticipantQuery.EXPECT().FindParticipantByID(gomock.Any(), gomock.Eq("user-1")).Return(mockParticipantEntity, nil).Times(1)

	service := NewParticipantQueryService(mockParticipantQuery)

	participant, err := service.GetParticipantByID(ctx, "user-1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if participant == nil {
		t.Fatal("expected participant, got nil")
	}

	// 期待するDTOを作成
	expectedDTO := dto.ParticipantEntityToResponse(mockParticipantEntity)

	// 詳細検証: 返されたDTOが期待するDTOと一致するか
	if !reflect.DeepEqual(participant, expectedDTO) {
		t.Errorf("returned participant does not match expected participant.\nExpected: %+v\nGot: %+v", expectedDTO, participant)
	}

	// Test case for not found
	mockParticipantQuery.EXPECT().FindParticipantByID(gomock.Any(), gomock.Eq("non-existent")).Return(nil, nil).Times(1) // nil entity, nil error を返す

	participantNotFound, errNotFound := service.GetParticipantByID(ctx, "non-existent")
	if errNotFound != nil {
		t.Fatalf("expected no error for not found, got %v", errNotFound)
	}
	if participantNotFound != nil {
		t.Errorf("expected nil participant for not found, got %+v", participantNotFound)
	}
}

func TestGetCoachesBySport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockParticipantQuery := mocks.NewMockParticipantQueryInterface(ctrl)
	ctx := context.Background()

	// participantQuery.FindCoachesBySport のモック設定
	mockCoachEntities := []entity.Participant{
		{ID: "coach-1", Name: "Coach A", Email: "coachA@example.com", Role: "coach", Sports: []string{"soccer"}},
		{ID: "coach-2", Name: "Coach B", Email: "coachB@example.com", Role: "coach", Sports: []string{"basketball"}},
	}
	mockParticipantQuery.EXPECT().FindCoachesBySport(gomock.Any(), gomock.Eq("soccer")).Return(mockCoachEntities, nil).Times(1)

	service := NewParticipantQueryService(mockParticipantQuery)

	coaches, err := service.GetCoachesBySport(ctx, "soccer")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(coaches) != 2 {
		t.Fatalf("expected 2 coaches, got %d", len(coaches))
	}

	// 期待するDTOのリストを作成
	expectedDTOs := []dto.ParticipantResponse{
		*dto.ParticipantEntityToResponse(&mockCoachEntities[0]),
		*dto.ParticipantEntityToResponse(&mockCoachEntities[1]),
	}

	// 詳細検証: 返されたDTOのリストが期待するリストと一致するか
	if !reflect.DeepEqual(coaches, expectedDTOs) {
		t.Errorf("returned coaches do not match expected coaches.\nExpected: %+v\nGot: %+v", expectedDTOs, coaches)
	}
}
