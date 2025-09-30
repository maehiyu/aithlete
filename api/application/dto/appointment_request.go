package dto

import "time"


type AppointmentCreateRequest struct {
	ChatID         string   `json:"chat_id" validate:"required"`
	Title          string   `json:"title" validate:"required"`
	Description    string   `json:"description"`
	ScheduledAt    time.Time `json:"scheduled_at" validate:"required"`
	Duration       int      `json:"duration" validate:"required,gt=0"`
	ParticipantIDs []string `json:"participant_ids" validate:"required,min=1"`
}


type AppointmentUpdateRequest struct {
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
	Status      *string    `json:"status,omitempty"`
}
