package query

import (
	"api/application/dto"
	"api/application/query"
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type appointmentQuery struct {
	db *pgxpool.Pool
}

func NewAppointmentQuery(db *pgxpool.Pool) query.AppointmentQueryInterface {
	return &appointmentQuery{db: db}
}

const baseAppointmentQuery = `
SELECT
    a.id, a.chat_id, a.title, a.description, a.scheduled_at, a.duration, a.status,
    COALESCE(
        (SELECT json_agg(json_build_object(
            'participant_id', ap.participant_id,
            'participation_status', ap.status
        ))
        FROM appointment_participants AS ap
        WHERE ap.appointment_id = a.id),
        '[]'::json
    ) AS participants
FROM
    appointments AS a
`

func (q *appointmentQuery) GetByID(ctx context.Context, id string) (*dto.AppointmentResponse, error) {
	query := baseAppointmentQuery + "WHERE a.id = $1"
	return q.scanOne(ctx, query, id)
}

func (q *appointmentQuery) ListByChatID(ctx context.Context, chatID string) ([]*dto.AppointmentResponse, error) {
	query := baseAppointmentQuery + "WHERE a.chat_id = $1 ORDER BY a.scheduled_at DESC"
	return q.scanList(ctx, query, chatID)
}

func (q *appointmentQuery) ListByUserID(ctx context.Context, userID string) ([]*dto.AppointmentResponse, error) {
	query := baseAppointmentQuery + `
        WHERE a.id IN (
            SELECT appointment_id FROM appointment_participants WHERE participant_id = $1
        )
        ORDER BY a.scheduled_at DESC
    `
	return q.scanList(ctx, query, userID)
}

func (q *appointmentQuery) ListByCoachID(ctx context.Context, coachID string) ([]*dto.AppointmentResponse, error) {
    // ListByUserIDと同じロジックで実装可能
	query := baseAppointmentQuery + `
        WHERE a.id IN (
            SELECT appointment_id FROM appointment_participants WHERE participant_id = $1
        )
        ORDER BY a.scheduled_at DESC
    `
	return q.scanList(ctx, query, coachID)
}

func (q *appointmentQuery) scanOne(ctx context.Context, query string, args ...interface{}) (*dto.AppointmentResponse, error) {
	row := q.db.QueryRow(ctx, query, args...)
	return q.scanRow(row)
}

func (q *appointmentQuery) scanList(ctx context.Context, query string, args ...interface{}) ([]*dto.AppointmentResponse, error) {
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	appointments := make([]*dto.AppointmentResponse, 0)
	for rows.Next() {
		appointment, err := q.scanRow(rows)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (q *appointmentQuery) scanRow(row pgx.Row) (*dto.AppointmentResponse, error) {
	var app dto.AppointmentResponse
	var participantsJSON []byte

	err := row.Scan(
		&app.ID,
		&app.ChatID,
		&app.Title,
		&app.Description,
		&app.ScheduledAt,
		&app.Duration,
		&app.Status,
		&participantsJSON,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil // Not found
		}
		log.Printf("Error scanning appointment row: %v", err) // デバッグログ追加
		return nil, err
	}

	if err := json.Unmarshal(participantsJSON, &app.Participants); err != nil {
		log.Printf("Error unmarshalling participants JSON: %v, JSON: %s", err, string(participantsJSON)) // デバッグログ追加
		return nil, err
	}

	return &app, nil
}