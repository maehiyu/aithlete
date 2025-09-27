package entity

import (
	"time"
)

type Appointment struct {
	ID          string
	ChatID      string
	CoachID     string
	UserID      string
	Title       string
	Description string
	ScheduledAt time.Time
	Duration    int
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
