package dto

import "time"

type AppointmentResponse struct {
	ID          string    `json:"id"`
	ChatID      string    `json:"chatId"`
	CoachID     string    `json:"coachId"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ScheduledAt time.Time `json:"scheduledAt"`
	Duration    int       `json:"duration"`
	Status      string    `json:"status"`
}
