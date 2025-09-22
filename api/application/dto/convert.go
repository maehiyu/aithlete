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
			Type:         "question",
		})
	}
	for _, a := range chat.Answers {
		chatItems = append(chatItems, ChatItem{
			ID:            a.ID,
			ParticipantID: a.ParticipantID,
			Content:       a.Content,
			CreatedAt:     a.CreatedAt,
			Type:         "answer",
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
			Role:    p.Role,
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

func QuestionsEntityToResponse(questions []entity.Question) []QuestionResponse {
	res := make([]QuestionResponse, len(questions))
	for i, q := range questions {
		res[i] = QuestionResponse{
			ID:            q.ID,
			ParticipantID: q.ParticipantID,
			Content:       q.Content,
			CreatedAt:     q.CreatedAt,
			Attachments:   []AttachmentResponse{}, // Add mapping if needed
		}
	}
	return res
}

func AnswersEntityToResponse(answers []entity.Answer) []AnswerResponse {
	res := make([]AnswerResponse, len(answers))
	for i, a := range answers {
		res[i] = AnswerResponse{
			ID:            a.ID,
			QuestionID:    a.QuestionID,
			ParticipantID: a.ParticipantID,
			Content:       a.Content,
			CreatedAt:     a.CreatedAt,
			Attachments:   []AttachmentResponse{}, // Add mapping if needed
		}
	}
	return res
}

func ParticipantEntityToResponse(p *entity.Participant) *ParticipantResponse {
	if p == nil {
		return nil
	}
	return &ParticipantResponse{
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

func QuestionEntityToResponse(q *entity.Question) QuestionResponse {
	return QuestionResponse{
		ID:            q.ID,
		ChatID:        q.ChatID,
		ParticipantID: q.ParticipantID,
		Content:       q.Content,
		CreatedAt:     q.CreatedAt,
		Attachments:   []AttachmentResponse{}, // Add mapping if needed
	}
}

func AnswerEntityToResponse(a *entity.Answer) AnswerResponse {
	return AnswerResponse{
		ID:            a.ID,
		ChatID:        a.ChatID,
		QuestionID:    a.QuestionID,
		ParticipantID: a.ParticipantID,
		Content:       a.Content,
		CreatedAt:     a.CreatedAt,
		Attachments:   []AttachmentResponse{}, // Add mapping if needed
	}
}

func QuestionResponseToEntity(resp QuestionResponse) *entity.Question {
	return &entity.Question{
		ID:            resp.ID,
		ChatID:        resp.ChatID,
		ParticipantID: resp.ParticipantID,
		Content:       resp.Content,
		CreatedAt:     resp.CreatedAt,
		// Attachments:   ... // 必要なら変換追加
	}
}

func AnswerResponseToEntity(resp AnswerResponse) *entity.Answer {
	return &entity.Answer{
		ID:            resp.ID,
		ChatID:        resp.ChatID,
		QuestionID:    resp.QuestionID,
		ParticipantID: resp.ParticipantID,
		Content:       resp.Content,
		CreatedAt:     resp.CreatedAt,
		// Attachments:   ... // 必要なら変換追加
	}
}