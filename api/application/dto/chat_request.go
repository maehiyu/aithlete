package dto

type ChatCreateRequest struct {
	Title          *string  `json:"title"`
	ParticipantIDs []string `json:"participant_ids"`
}

type ChatUpdateRequest struct {
	Title          *string  `json:"title"`
	ParticipantIDs []string `json:"participant_ids"`
}

type ParticipantCreateRequest struct {
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Role   string   `json:"role"`
	Sports []string `json:"sports"`
	IconURL *string  `json:"icon_url"`
}

type ParticipantUpdateRequest struct {
	Name   *string   `json:"name"`
	Email  *string   `json:"email"`
	Role   *string   `json:"role"`
	Sports []string `json:"sports"`
	IconURL *string  `json:"icon_url"`
}

type QuestionCreateRequest struct {
	ParticipantID string
	Content       string
}

type AnswerCreateRequest struct {
	QuestionID    string
	ParticipantID string
	Content       string
}