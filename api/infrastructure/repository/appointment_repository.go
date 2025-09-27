package repository

import (
	"api/domain/entity"
	"api/domain/repository"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type appointmentRepository struct {
	db *pgxpool.Pool
}

func NewAppointmentRepository(db *pgxpool.Pool) repository.AppointmentRepositoryInterface {
	return &appointmentRepository{db: db}
}

func (r *appointmentRepository) Create(ctx context.Context, appointment *entity.Appointment) (string, error) {
	query := `
		INSERT INTO appointments (id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id
	`
	var id string
	err := r.db.QueryRow(ctx, query,
		appointment.ID,
		appointment.ChatID,
		appointment.CoachID,
		appointment.UserID,
		appointment.Title,
		appointment.Description,
		appointment.ScheduledAt,
		appointment.Duration,
		appointment.Status,
		appointment.CreatedAt,
		appointment.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *appointmentRepository) FindByID(ctx context.Context, id string) (*entity.Appointment, error) {
	query := `SELECT id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE id = $1`
	appointment := &entity.Appointment{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&appointment.ID,
		&appointment.ChatID,
		&appointment.CoachID,
		&appointment.UserID,
		&appointment.Title,
		&appointment.Description,
		&appointment.ScheduledAt,
		&appointment.Duration,
		&appointment.Status,
		&appointment.CreatedAt,
		&appointment.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // 見つからない場合はエラーではなくnilを返す
		}
		return nil, err
	}
	return appointment, nil
}

func (r *appointmentRepository) Update(ctx context.Context, appointment *entity.Appointment) error {
	query := `
		UPDATE appointments
		SET title = $2, description = $3, scheduled_at = $4, duration = $5, status = $6, updated_at = $7
		WHERE id = $1
	`
	_, err := r.db.Exec(ctx, query,
		appointment.ID,
		appointment.Title,
		appointment.Description,
		appointment.ScheduledAt,
		appointment.Duration,
		appointment.Status,
		appointment.UpdatedAt,
	)
	return err
}

func (r *appointmentRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

var _ repository.AppointmentRepositoryInterface = (*appointmentRepository)(nil)