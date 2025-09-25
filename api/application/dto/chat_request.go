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
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Role    string   `json:"role"`
	Sports  []string `json:"sports"`
	IconURL *string  `json:"icon_url"`
}

type ParticipantUpdateRequest struct {
	Name    *string  `json:"name"`
	Email   *string  `json:"email"`
	Role    *string  `json:"role"`
	Sports  []string `json:"sports"`
	IconURL *string  `json:"icon_url"`
}

type QuestionCreateRequest struct {
	ParticipantID string `json:"participant_id"`
	Content       string `json:"content"`
}

type AnswerCreateRequest struct {
	QuestionID    string `json:"question_id"`
	ParticipantID string `json:"participant_id"`
	Content       string `json:"content"`
}
type ChatItemRequest struct {
	ParticipantID string `json:"participant_id"`
	QuestionID   *string `json:"question_id,omitempty"` 
	Content       string `json:"content"`
	Type          string `json:"type"` // "question", "answer", "ai_answer"
	TempID        string `json:"temp_id"`
}