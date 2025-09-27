//go:generate mockgen -source=appointment_repository.go -destination=mocks/mock_appointment_repository.go -package=mocks AppointmentRepositoryInterface
package repository

import (
	"api/domain/entity"
	"context"
)

type AppointmentRepositoryInterface interface {
	Create(ctx context.Context, appointment *entity.Appointment) (string, error)
	FindByID(ctx context.Context, id string) (*entity.Appointment, error)
	Update(ctx context.Context, appointment *entity.Appointment) error
	Delete(ctx context.Context, id string) error
}

