package query

import (
	"api/application/dto"
	"api/application/query"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ParticipantQuery struct {
	conn *pgxpool.Pool
}

func NewParticipantQuery(conn *pgxpool.Pool) *ParticipantQuery {
	return &ParticipantQuery{conn: conn}
}

// FindParticipantsByChatID: チャットIDに紐づく参加者一覧を取得
func (q *ParticipantQuery) FindParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error) {
	ctx := context.Background()
	row := q.conn.QueryRow(ctx, `SELECT participant_ids FROM chats WHERE id = $1`, chatID)
	var participantIDs []string
	if err := row.Scan(&participantIDs); err != nil {
		return nil, err
	}
	participants := []dto.ParticipantResponse{}
	for _, pid := range participantIDs {
		prow := q.conn.QueryRow(ctx, `SELECT id, name, email, role, icon_url, sports FROM participants WHERE id = $1`, pid)
		var p dto.ParticipantResponse
		var iconURL *string
		if err := prow.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &iconURL, &p.Sports); err == nil {
			p.IconURL = iconURL
			participants = append(participants, p)
		}
	}
	return participants, nil
}

// FindParticipantByID: 参加者IDで参加者情報を取得
func (q *ParticipantQuery) FindParticipantByID(participantID string) (*dto.ParticipantResponse, error) {
	ctx := context.Background()
	row := q.conn.QueryRow(ctx, `SELECT id, name, email, role, icon_url, sports FROM participants WHERE id = $1`, participantID)
	var p dto.ParticipantResponse
	var iconURL *string
	if err := row.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &iconURL, &p.Sports); err != nil {
		return nil, err
	}
	p.IconURL = iconURL
	return &p, nil
}

var _ interface {
	query.ParticipantQueryInterface
} = (*ParticipantQuery)(nil)
