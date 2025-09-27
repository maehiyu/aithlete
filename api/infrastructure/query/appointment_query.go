package query

import (
	"api/application/dto"
	"api/application/query"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type appointmentQuery struct {
	db *pgxpool.Pool
}

func NewAppointmentQuery(db *pgxpool.Pool) query.AppointmentQueryInterface {
	return &appointmentQuery{db: db}
}

func (q *appointmentQuery) GetByID(ctx context.Context, id string) (*dto.AppointmentResponse, error) {
	query := `SELECT id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE id = $1`
	appointment := &dto.AppointmentResponse{}
	err := q.db.QueryRow(ctx, query, id).Scan(
		&appointment.ID,
		&appointment.ChatID,
		&appointment.CoachID,
		&appointment.UserID,
		&appointment.Title,
		&appointment.Description,
		&appointment.ScheduledAt,
		&appointment.Duration,
		&appointment.Status,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // 見つからない場合はエラーではなくnilを返す
		}
		return nil, err
	}
	return appointment, nil
}

func (q *appointmentQuery) GetByChatID(ctx context.Context, chatID string) ([]*dto.AppointmentResponse, error) {
	query := `SELECT id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE chat_id = $1 ORDER BY scheduled_at DESC`
	return q.queryAppointments(ctx, query, chatID)
}

func (q *appointmentQuery) GetByUserID(ctx context.Context, userID string) ([]*dto.AppointmentResponse, error) {
	query := `SELECT id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE user_id = $1 ORDER BY scheduled_at DESC`
	return q.queryAppointments(ctx, query, userID)
}

func (q *appointmentQuery) GetByCoachID(ctx context.Context, coachID string) ([]*dto.AppointmentResponse, error) {
	query := `SELECT id, chat_id, coach_id, user_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE coach_id = $1 ORDER BY scheduled_at DESC`
	return q.queryAppointments(ctx, query, coachID)
}

// queryAppointments は複数件取得の共通ロジックです。
func (q *appointmentQuery) queryAppointments(ctx context.Context, query string, args ...interface{}) ([]*dto.AppointmentResponse, error) {
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]*dto.AppointmentResponse, 0)
	for rows.Next() {
		appointment := &dto.AppointmentResponse{}
		if err := rows.Scan(
			&appointment.ID,
			&appointment.ChatID,
			&appointment.CoachID,
			&appointment.UserID,
			&appointment.Title,
			&appointment.Description,
			&appointment.ScheduledAt,
			&appointment.Duration,
			&appointment.Status,
		); err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}
