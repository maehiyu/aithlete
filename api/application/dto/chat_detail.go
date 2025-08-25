package dto

import "time"

type ChatDetailResponse struct {
	ID           string                   `json:"id"`
	Title        *string                  `json:"title,omitempty"`
	Participants []ParticipantResponse    `json:"participants"`
	Questions    []QuestionDetailResponse `json:"questions"`
	Answers      []AnswerDetailResponse   `json:"answers"`
	StartedAt    time.Time                `json:"started_at"`
	LastActiveAt time.Time                `json:"last_active_at"`
}

type ParticipantResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	IconURL *string `json:"icon_url"`
	Sports  []string `json:"sports"`
}

type QuestionDetailResponse struct {
	ID            string               `json:"id"`
	ParticipantID string               `json:"participant_id"`
	Content       string               `json:"content"`
	Attachments   []AttachmentResponse `json:"attachments"`
	CreatedAt     time.Time            `json:"created_at"`
}

type AnswerDetailResponse struct {
	ID            string               `json:"id"`
	QuestionID    string               `json:"question_id"`
	ParticipantID string               `json:"participant_id"`
	Content       string               `json:"content"`
	Attachments   []AttachmentResponse `json:"attachments"`
	CreatedAt     time.Time            `json:"created_at"`
}

type AttachmentResponse struct {
	Type string            `json:"type"`
	URL  string            `json:"url"`
	Pose *PoseDataResponse `json:"pose,omitempty"`
}

type PoseDataResponse struct {
	Keypoints string  `json:"keypoints"`
	Score     float64 `json:"score"`
}
