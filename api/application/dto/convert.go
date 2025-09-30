package dto

import (
	"api/domain/entity"
	"sort"
	"time"
)

func ChatCreateRequestToEntity(dto ChatCreateRequest, id string, now time.Time) *entity.Chat {
	return &entity.Chat{
		ID:             id,
		Title:          dto.Title,
		ParticipantIDs: dto.ParticipantIDs,
		Questions:      []entity.Question{},
		Answers:        []entity.Answer{},
		StartedAt:      now,
		LastActiveAt:   now,
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

func ChatEntityToDetailResponse(chat *entity.Chat, participants []entity.Participant) *ChatDetailResponse {
	chatItems := make([]ChatItem, 0)
	for _, q := range chat.Questions {
		chatItems = append(chatItems, ChatItem{
			ID:            q.ID,
			ParticipantID: q.ParticipantID,
			Content:       q.Content,
			CreatedAt:     q.CreatedAt,
			Type:          "question",
		})
	}
	for _, a := range chat.Answers {
		chatItems = append(chatItems, ChatItem{
			ID:            a.ID,
			ParticipantID: a.ParticipantID,
			Content:       a.Content,
			CreatedAt:     a.CreatedAt,
			Type:          "answer",
		})
	}

	sort.Slice(chatItems, func(i, j int) bool {
		return chatItems[i].CreatedAt.Before(chatItems[j].CreatedAt)
	})

	participantResponses := make([]ParticipantResponse, len(participants))
	for i, p := range participants {
		participantResponses[i] = ParticipantResponse{
			ID:      p.ID,
			Name:    p.Name,
			Email:   p.Email,
			Role:    p.Role,
			Sports:  p.Sports,
			IconURL: p.IconURL,
		}
	}

	return &ChatDetailResponse{
		ID:           chat.ID,
		Title:        chat.Title,
		Participants: participantResponses,
		Timeline:     chatItems,
		StartedAt:    chat.StartedAt,
		LastActiveAt: chat.LastActiveAt,
	}
}

func ParticipantEntityToResponse(p *entity.Participant) *ParticipantResponse {
	if p == nil {
		return nil
	}
	return &ParticipantResponse{
		ID:      p.ID,
		Name:    p.Name,
		Email:   p.Email,
		Role:    p.Role,
		Sports:  p.Sports,
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

func QuestionCreateRequestToEntity(dto QuestionCreateRequest, id string, chatId string, now time.Time) *entity.Question {
	return &entity.Question{
		ID:            id,
		ChatID:        chatId,
		ParticipantID: dto.ParticipantID,
		Content:       dto.Content,
		CreatedAt:     now,
	}
}

func AnswerCreateRequestToEntity(dto AnswerCreateRequest, id string, chatId string, questionID string, now time.Time) *entity.Answer {
	return &entity.Answer{
		ID:            id,
		ChatID:        chatId,
		QuestionID:    questionID,
		ParticipantID: dto.ParticipantID,
		Content:       dto.Content,
		CreatedAt:     now,
	}
}

func ChatItemToQuestion(resp ChatItem) *entity.Question {
	return &entity.Question{
		ID:            resp.ID,
		ChatID:        resp.ChatID,
		ParticipantID: resp.ParticipantID,
		Content:       resp.Content,
		CreatedAt:     resp.CreatedAt,
		// Attachments:   ... // 必要なら変換追加
	}
}

func ChatItemToAnswer(resp ChatItem) *entity.Answer {
	var questionID string
	if resp.QuestionID != nil {
		questionID = *resp.QuestionID
	}
	return &entity.Answer{
		ID:            resp.ID,
		ChatID:        resp.ChatID,
		QuestionID:    questionID,
		ParticipantID: resp.ParticipantID,
		Content:       resp.Content,
		CreatedAt:     resp.CreatedAt,
		// Attachments:   ... // 必要なら変換追加
	}
}

func AppointmentCreateRequestToEntity(dto AppointmentCreateRequest, id string, now time.Time) *entity.Appointment {
	return &entity.Appointment{
		ID:          id,
		ChatID:      dto.ChatID,
		Title:       dto.Title,
		Description: dto.Description,
		ScheduledAt: dto.ScheduledAt,
		Duration:    dto.Duration,
		Status:      "scheduled",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func AppointmentUpdateRequestToEntity(appointment *entity.Appointment, dto AppointmentUpdateRequest) {
	if dto.Title != nil {
		appointment.Title = *dto.Title
	}
	if dto.Description != nil {
		appointment.Description = *dto.Description
	}
	if dto.ScheduledAt != nil {
		appointment.ScheduledAt = *dto.ScheduledAt
	}
	if dto.Duration != nil {
		appointment.Duration = *dto.Duration
	}
	if dto.Status != nil {
		appointment.Status = *dto.Status
	}
	appointment.UpdatedAt = time.Now()
}

func AppointmentEntityToResponses(appointment *entity.Appointment) *AppointmentResponse {
	if appointment == nil {
		return nil
	}
	return &AppointmentResponse{
		ID:          appointment.ID,
		ChatID:      appointment.ChatID,
		Title:       appointment.Title,
		Description: appointment.Description,
		ScheduledAt: appointment.ScheduledAt,
		Duration:    appointment.Duration,
		Status:      appointment.Status,
	}
}
