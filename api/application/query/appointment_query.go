//go:generate mockgen -source=appointment_query.go -destination=mocks/mock_appointment_query.go -package=mocks AppointmentQueryInterface
package query

import (
	"api/application/dto"
	"context"
)

type AppointmentQueryInterface interface {
	GetByID(ctx context.Context, id string) (*dto.AppointmentResponse, error)
	GetByChatID(ctx context.Context, chatID string) ([]*dto.AppointmentResponse, error)
	GetByUserID(ctx context.Context, userID string) ([]*dto.AppointmentResponse, error)
	GetByCoachID(ctx context.Context, coachID string) ([]*dto.AppointmentResponse, error)
}