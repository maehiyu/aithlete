package dto

import (
	"api/domain/entity"
	"time"
)

func ChatCreateRequestToEntity(dto ChatCreateRequest, id string, now time.Time) *entity.Chat {
	return &entity.Chat{
		ID:            id,
		Title:         dto.Title,
		ParticipantIDs: dto.ParticipantIDs,
		Questions:     []entity.Question{},
		Answers:       []entity.Answer{},
		StartedAt:     now,
		LastActiveAt:  now,
	}
}

func ChatUpdateRequestToEntity(chat *entity.Chat, dto ChatUpdateRequest) {
	if dto.Title != nil {
		chat.Title = dto.Title
	}
	if dto.ParticipantIDs != nil {
		chat.ParticipantIDs = dto.ParticipantIDs
	}
}

func ChatEntityToDetailResponse(chat *entity.Chat, participants []ParticipantResponse) ChatDetailResponse {
	return ChatDetailResponse{
		ID:           chat.ID,
		Title:        chat.Title,
		Participants: participants,
		Questions:    []QuestionDetailResponse{},
		Answers:      []AnswerDetailResponse{},
		StartedAt:    chat.StartedAt,
		LastActiveAt: chat.LastActiveAt,
	}
}

func ParticipantEntityToResponse(p *entity.Participant) ParticipantResponse {
	return ParticipantResponse{
		ID:      p.ID,
		Name:    p.Name,
		Role:    p.Role,
		IconURL: p.IconURL,
	}
}

func ParticipantCreateRequestToEntity(dto ParticipantCreateRequest, id string) *entity.Participant {
	return &entity.Participant{
		ID:      id,
		Name:    dto.Name,
		Role:    dto.Role,
		Sports:  dto.Sports,
		IconURL: dto.IconURL,
	}
}

func ParticipantUpdateRequestToEntity(participant *entity.Participant, dto ParticipantUpdateRequest) {
	if dto.Name != nil {
		participant.Name = *dto.Name
	}
	if dto.Role != nil {
		participant.Role = *dto.Role
	}
	if dto.Sports != nil {
		participant.Sports = dto.Sports
	}
	if dto.IconURL != nil {
		participant.IconURL = dto.IconURL
	}
}