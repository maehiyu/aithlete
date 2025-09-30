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

// --- Appointment（予約枠）の操作 ---

func (r *appointmentRepository) CreateAppointment(ctx context.Context, appointment *entity.Appointment) (*entity.Appointment, error) {
	query := `
		INSERT INTO appointments (id, chat_id, title, description, scheduled_at, duration, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, chat_id, title, description, scheduled_at, duration, status, created_at, updated_at
	`
	created := &entity.Appointment{}
	err := r.db.QueryRow(ctx, query,
		appointment.ID,
		appointment.ChatID,
		appointment.Title,
		appointment.Description,
		appointment.ScheduledAt,
		appointment.Duration,
		appointment.Status,
		appointment.CreatedAt,
		appointment.UpdatedAt,
	).Scan(
		&created.ID,
		&created.ChatID,
		&created.Title,
		&created.Description,
		&created.ScheduledAt,
		&created.Duration,
		&created.Status,
		&created.CreatedAt,
		&created.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *appointmentRepository) FindAppointmentByID(ctx context.Context, id string) (*entity.Appointment, error) {
	query := `SELECT id, chat_id, title, description, scheduled_at, duration, status, created_at, updated_at FROM appointments WHERE id = $1`
	appointment := &entity.Appointment{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&appointment.ID,
		&appointment.ChatID,
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

func (r *appointmentRepository) UpdateAppointment(ctx context.Context, appointment *entity.Appointment) error {
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

func (r *appointmentRepository) DeleteAppointment(ctx context.Context, id string) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// --- AppointmentParticipant（参加情報）の操作 ---

func (r *appointmentRepository) CreateAppointmentParticipants(ctx context.Context, participants []*entity.AppointmentParticipant) error {
	rows := make([][]interface{}, len(participants))
	for i, p := range participants {
		rows[i] = []interface{}{p.AppointmentID, p.ParticipantID, p.Status}
	}

	_, err := r.db.CopyFrom(
		ctx,
		pgx.Identifier{"appointment_participants"},
		[]string{"appointment_id", "participant_id", "status"},
		pgx.CopyFromRows(rows),
	)
	return err
}

func (r *appointmentRepository) UpdateAppointmentParticipant(ctx context.Context, participant *entity.AppointmentParticipant) error {
	query := `
		UPDATE appointment_participants
		SET status = $3
		WHERE appointment_id = $1 AND participant_id = $2
	`
	_, err := r.db.Exec(ctx, query, participant.AppointmentID, participant.ParticipantID, participant.Status)
	return err
}

func (r *appointmentRepository) RemoveAppointmentParticipant(ctx context.Context, appointmentID string, participantID string) error {
	query := `DELETE FROM appointment_participants WHERE appointment_id = $1 AND participant_id = $2`
	_, err := r.db.Exec(ctx, query, appointmentID, participantID)
	return err
}

func (r *appointmentRepository) FindParticipantsByAppointmentID(ctx context.Context, appointmentID string) ([]*entity.AppointmentParticipant, error) {
	query := `SELECT appointment_id, participant_id, status FROM appointment_participants WHERE appointment_id = $1`
	rows, err := r.db.Query(ctx, query, appointmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	participants := make([]*entity.AppointmentParticipant, 0)
	for rows.Next() {
		participant := &entity.AppointmentParticipant{}
		if err := rows.Scan(&participant.AppointmentID, &participant.ParticipantID, &participant.Status); err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	return participants, nil
}

var _ repository.AppointmentRepositoryInterface = (*appointmentRepository)(nil)
