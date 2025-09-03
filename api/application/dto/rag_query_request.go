package dto

type RagQueryRequest struct {
	Question      string `json:"question"`
	ChatID        string `json:"chat_id"`
	QuestionID    string `json:"question_id"`
	ParticipantID string `json:"participant_id"`
}