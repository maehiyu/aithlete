package query

import (
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
	rows, err := q.pool.Query(ctx, "SELECT id, name, email FROM participants WHERE id = ANY($1)", ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []entity.Participant
	for rows.Next() {
		var p entity.Participant
		if err := rows.Scan(&p.ID, &p.Name, &p.Email); err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}