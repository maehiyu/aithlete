package dto

import "time"

type ChatDetailResponse struct {
	ID           string                `json:"id"`
	Title        *string               `json:"title"`
	StartedAt    time.Time             `json:"started_at"`
	LastActiveAt time.Time             `json:"last_active_at"`
	Participants []ParticipantResponse `json:"participants"`
	Questions    []QuestionResponse    `json:"questions"`
	Answers      []AnswerResponse      `json:"answers"`
}

type ChatSummaryResponse struct {
	ID           string    `json:"id"`
	Title        *string   `json:"title"`
	LastActiveAt time.Time `json:"last_active_at"`
	LatestQA	 *string   `json:"latest_qa"`
	Opponent   	 OpponentResponse `json:"opponent"`
}

type ParticipantResponse struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Email  string   `json:"email"`
	Role   string   `json:"role"`
	Sports []string `json:"sports"`
	IconURL *string  `json:"icon_url"`
}

type OpponentResponse struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Role   string   `json:"role"`
	IconURL *string  `json:"icon_url"`
}

type QuestionResponse struct {
	ID            string    `json:"id"`
	ParticipantID string    `json:"participant_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	Attachments   []AttachmentResponse `json:"attachments"`
}

type AnswerResponse struct {
	ID            string    `json:"id"`
	QuestionID    string    `json:"question_id"`
	ParticipantID string    `json:"participant_id"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	Attachments   []AttachmentResponse `json:"attachments"`
}

type AttachmentResponse struct {
	ID   string  `json:"id"`
	Type string  `json:"type"`
	URL  *string `json:"url"`
}