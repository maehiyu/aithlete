package command

import (
	"api/application/dto"
	"api/domain/entity"
	"testing"
)

func TestCreateParticipant_Success(t *testing.T) {
	mockRepo := NewMockParticipantRepository()
	mockRepo.CreateFunc = func(p *entity.Participant) (*entity.Participant, error) {
		return p, nil
	}

	svc := &ParticipantCommandService{
		participantRepo: mockRepo,
	}

	name := "test user"
	role := "user"
	iconURL := "https://example.com/icon.png"
	createReq := dto.ParticipantCreateRequest{
		Name:    name,
		Role:    role,
		IconURL: &iconURL,
	}

	resp, err := svc.CreateParticipant(createReq, "test-participant-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Name != name {
		t.Errorf("expected name %v, got %v", name, resp.Name)
	}
	if resp.Role != role {
		t.Errorf("expected role %v, got %v", role, resp.Role)
	}
	if resp.IconURL == nil || *resp.IconURL != iconURL {
		t.Errorf("expected iconURL %v, got %v", iconURL, resp.IconURL)
	}
}

func TestUpdateParticipant_Success(t *testing.T) {
	mockRepo := NewMockParticipantRepository()
	// 既存の参加者
	before := &entity.Participant{
		ID:      "1",
		Name:    "old name",
		Role:    "user",
		IconURL: nil,
	}
	// FindByIDで既存参加者を返す
	mockRepo.FindByIDFunc = func(id string) (*entity.Participant, error) {
		return before, nil
	}
	// Updateはそのまま返す
	mockRepo.UpdateFunc = func(p *entity.Participant) (*entity.Participant, error) {
		return p, nil
	}

	svc := &ParticipantCommandService{
		participantRepo: mockRepo,
	}

	newName := "new name"
	newRole := "coach"
	newIcon := "https://example.com/newicon.png"
	updateReq := dto.ParticipantUpdateRequest{
		Name:    &newName,
		Role:    &newRole,
		IconURL: &newIcon,
	}

	resp, err := svc.UpdateParticipant("1",updateReq)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if resp.Name != newName {
		t.Errorf("expected name %v, got %v", newName, resp.Name)
	}
	if resp.Role != newRole {
		t.Errorf("expected role %v, got %v", newRole, resp.Role)
	}
	if resp.IconURL == nil || *resp.IconURL != newIcon {
		t.Errorf("expected iconURL %v, got %v", newIcon, resp.IconURL)
	}
}
