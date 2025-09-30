package dto

import "time"

// AppointmentParticipantResponse は予約レスポンスに含まれる参加者の状態を表します。
type AppointmentParticipantResponse struct {
	ParticipantID     string `json:"participant_id"`
	ParticipationStatus string `json:"participation_status"` // "needs-action", "accepted", "declined", "tentative"
}

// AppointmentResponse はクライアントに返す予約情報を表します。
type AppointmentResponse struct {
	ID           string                           `json:"id"`
	ChatID       string                           `json:"chat_id"`
	Title        string                           `json:"title"`
	Description  string                           `json:"description"`
	ScheduledAt  time.Time                        `json:"scheduled_at"`
	Duration     int                              `json:"duration"`
	Status       string                           `json:"status"`
	Participants []AppointmentParticipantResponse `json:"participants"`
}
