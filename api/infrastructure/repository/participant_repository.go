// participants テーブル定義
// id TEXT PRIMARY KEY
// name TEXT NOT NULL
// email TEXT NOT NULL
// role TEXT NOT NULL
// sports TEXT[] NOT NULL
// icon_url TEXT
package repository

import (
	"api/domain/entity"
	"api/domain/repository"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ParticipantRepositoryImpl struct {
	conn *pgxpool.Pool
}

func NewParticipantRepository(conn *pgxpool.Pool) *ParticipantRepositoryImpl {
	return &ParticipantRepositoryImpl{conn: conn}
}

func (r *ParticipantRepositoryImpl) FindByID(ctx context.Context, participantID string) (*entity.Participant, error) {
	var p entity.Participant
	row := r.conn.QueryRow(ctx, `SELECT id, name, email, role, sports, icon_url FROM participants WHERE id = $1`, participantID)
	err := row.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ParticipantRepositoryImpl) FindByIDs(ctx context.Context, participantIDs []string) ([]*entity.Participant, error) {
	rows, err := r.conn.Query(ctx, `SELECT id, name, email, role, sports, icon_url FROM participants WHERE id = ANY($1)`, participantIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []*entity.Participant
	for rows.Next() {
		var p entity.Participant
		err := rows.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &p.Sports, &p.IconURL)
		if err != nil {
			return nil, err
		}
		participants = append(participants, &p)
	}
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	return participants, nil
}

func (r *ParticipantRepositoryImpl) Create(ctx context.Context, participant *entity.Participant) (string, error) {
	log.Printf("Creating participant with sports: %v", participant.Sports)
	_, err := r.conn.Exec(ctx,
		`INSERT INTO participants (id, name, email, role, sports, icon_url) VALUES ($1, $2, $3, $4, $5, $6)`,
		participant.ID, participant.Name, participant.Email, participant.Role, participant.Sports, participant.IconURL,
	)
	if err != nil {
		log.Printf("Error creating participant: %v", err)
		return "", err
	}
	return participant.ID, nil
}

func (r *ParticipantRepositoryImpl) Update(ctx context.Context, participant *entity.Participant) error {
	log.Printf("Updating participant %s with sports: %v", participant.ID, participant.Sports)
	_, err := r.conn.Exec(ctx,
		`UPDATE participants SET name = $2, email = $3, role = $4, sports = $5, icon_url = $6 WHERE id = $1`,
		participant.ID, participant.Name, participant.Email, participant.Role, participant.Sports, participant.IconURL,
	)
	if err != nil {
		log.Printf("Error updating participant: %v", err)
		return err
	}
	log.Printf("Successfully updated participant %s", participant.ID)
	return nil
}

var _ repository.ParticipantRepositoryInterface = (*ParticipantRepositoryImpl)(nil)
