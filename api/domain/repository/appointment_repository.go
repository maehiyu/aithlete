//go:generate mockgen -source=appointment_repository.go -destination=mocks/mock_appointment_repository.go -package=mocks AppointmentRepositoryInterface
package repository

import (
	"api/domain/entity"
	"context"
)

type AppointmentRepositoryInterface interface {
	// Appointment（予約枠）の操作
	CreateAppointment(ctx context.Context, appointment *entity.Appointment) (*entity.Appointment, error)
	FindAppointmentByID(ctx context.Context, id string) (*entity.Appointment, error)
	UpdateAppointment(ctx context.Context, appointment *entity.Appointment) error
	DeleteAppointment(ctx context.Context, id string) error

	// AppointmentParticipant（参加情報）の操作
	CreateAppointmentParticipants(ctx context.Context, participants []*entity.AppointmentParticipant) error
	UpdateAppointmentParticipant(ctx context.Context, participant *entity.AppointmentParticipant) error
	RemoveAppointmentParticipant(ctx context.Context, appointmentID string, participantID string) error
	FindParticipantsByAppointmentID(ctx context.Context, appointmentID string) ([]*entity.AppointmentParticipant, error)
}

