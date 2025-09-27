package query

import (
	"api/application/query"
	"api/domain/entity"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ParticipantQuery struct {
	pool *pgxpool.Pool
}

func NewParticipantQuery(pool *pgxpool.Pool) *ParticipantQuery {
	return &ParticipantQuery{
		pool: pool,
	}
}

func (q *ParticipantQuery) FindParticipantsByIDs(ctx context.Context, ids []string) ([]entity.Participant, error) {
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, "SELECT id, name, email, role, sports, icon_url FROM participants WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []entity.Participant
	for rows.Next() {
		var p entity.Participant
		if err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL); err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

func (q *ParticipantQuery) FindParticipantsByChatID(ctx context.Context, chatID string) ([]entity.Participant, error) {
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `
		SELECT p.id, p.name, p.email, p.role, p.sports, p.icon_url
		FROM participants p
		JOIN chats c ON c.participant_ids @> ARRAY[p.id]
		WHERE c.id = $1
	`, chatID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []entity.Participant
	for rows.Next() {
		var p entity.Participant
		if err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL); err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

func (q *ParticipantQuery) FindParticipantByID(ctx context.Context, participantID string) (*entity.Participant, error) {
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, `
		SELECT id, name, email, role, sports, icon_url
		FROM participants WHERE id = $1
	`, participantID)

	var p entity.Participant
	if err := row.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL); err != nil {
		return nil, err
	}

	return &p, nil
}

func (q *ParticipantQuery) FindCoachesBySport(ctx context.Context, sport string) ([]entity.Participant, error) {
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `
		SELECT id, name, email, role, sports, icon_url
		FROM participants
		WHERE role = 'coach' AND $1 = ANY(sports)
	`, sport)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var coaches []entity.Participant
	for rows.Next() {
		var p entity.Participant
		if err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL); err != nil {
			return nil, err
		}
		coaches = append(coaches, p)
	}

	return coaches, nil
}

var _ query.ParticipantQueryInterface = (*ParticipantQuery)(nil)
