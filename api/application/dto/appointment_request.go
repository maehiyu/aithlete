package dto

import "time"

type AppointmentCreateRequest struct {
	ChatID      string    `json:"chatId" validate:"required"`
	CoachID     string    `json:"coachId" validate:"required"`
	UserID      string    `json:"userId" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	ScheduledAt time.Time `json:"scheduledAt" validate:"required"`
	Duration    int       `json:"duration" validate:"required,gt=0"`
}

type AppointmentUpdateRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	ScheduledAt *time.Time `json:"scheduledAt,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
	Status      *string    `json:"status,omitempty"`
}
