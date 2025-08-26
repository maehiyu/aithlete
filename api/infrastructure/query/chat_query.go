package query

import (
	"api/application/dto"
	"api/application/query"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatQuery struct {
	pool *pgxpool.Pool
}

func NewChatQuery(pool *pgxpool.Pool) *ChatQuery {
	return &ChatQuery{pool: pool}
}

// FindChatsByUserID: ユーザーが参加しているチャット一覧（サマリー）を取得
func (q *ChatQuery) FindChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error) {
	ctx := context.Background()
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	rows, err := conn.Query(ctx, `
		SELECT c.id, c.title, c.last_active_at,
			(
				SELECT content FROM (
					SELECT content, created_at FROM questions WHERE chat_id = c.id
					UNION ALL
					SELECT content, created_at FROM answers WHERE chat_id = c.id
				) qa
				ORDER BY created_at DESC LIMIT 1
			) AS latest_qa,
			p.id, p.name, p.role, p.icon_url
		FROM chats c
		LEFT JOIN LATERAL (
			SELECT id, name, role, icon_url
			FROM participants
			WHERE id = ANY(c.participant_ids) AND id <> $1
			LIMIT 1
		) p ON true
		WHERE $1 = ANY(c.participant_ids)
		ORDER BY c.last_active_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []dto.ChatSummaryResponse
	for rows.Next() {
		var chat dto.ChatSummaryResponse
		var title *string
		var latestQA *string
		var opponentID, opponentName, opponentRole, opponentIconURL *string
		if err := rows.Scan(&chat.ID, &title, &chat.LastActiveAt, &latestQA, &opponentID, &opponentName, &opponentRole, &opponentIconURL); err != nil {
			return nil, err
		}
		chat.Title = title
		chat.LatestQA = latestQA
		if opponentID != nil {
			chat.Opponent = dto.OpponentResponse{
				ID:      *opponentID,
				Name:    derefString(opponentName),
				Role:    derefString(opponentRole),
				IconURL: opponentIconURL,
			}
		}
		result = append(result, chat)
	}
	return result, nil
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// FindChatByID: チャット詳細を取得
func (q *ChatQuery) FindChatByID(chatID string) (*dto.ChatDetailResponse, error) {
	ctx := context.Background()
	conn, err := q.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	row := conn.QueryRow(ctx, `
		SELECT id, title, started_at, last_active_at, participant_ids
		FROM chats WHERE id = $1
	`, chatID)

	var chat dto.ChatDetailResponse
	var title *string
	var participantIDs []string
	if err := row.Scan(&chat.ID, &title, &chat.StartedAt, &chat.LastActiveAt, &participantIDs); err != nil {
		return nil, err
	}
	chat.Title = title

	participants := []dto.ParticipantResponse{}
	for _, pid := range participantIDs {
		prow := conn.QueryRow(ctx, `SELECT id, name, email, role, icon_url, sports FROM participants WHERE id = $1`, pid)
		var p dto.ParticipantResponse
		var iconURL *string
		if err := prow.Scan(&p.ID, &p.Name, &p.Email, &p.Role, &iconURL, &p.Sports); err == nil {
			p.IconURL = iconURL
			participants = append(participants, p)
		}
	}
	chat.Participants = participants

	qRows, err := conn.Query(ctx, `SELECT id, participant_id, content, created_at FROM questions WHERE chat_id = $1 ORDER BY created_at`, chatID)
	if err != nil {
		return nil, err
	}
	defer qRows.Close()
	questions := []dto.QuestionResponse{}
	for qRows.Next() {
		var qd dto.QuestionResponse
		if err := qRows.Scan(&qd.ID, &qd.ParticipantID, &qd.Content, &qd.CreatedAt); err != nil {
			return nil, err
		}

		// attachmentsは別コネクションで取得
		attConn, err := q.pool.Acquire(ctx)
		if err != nil {
			return nil, err
		}
		aRows, err := attConn.Query(ctx, `SELECT type, url FROM attachments WHERE question_id = $1`, qd.ID)
		if err != nil {
			attConn.Release()
			return nil, err
		}
		attachments := []dto.AttachmentResponse{}
		for aRows.Next() {
			var att dto.AttachmentResponse
			if err := aRows.Scan(&att.Type, &att.URL); err == nil {
				attachments = append(attachments, att)
			}
		}
		aRows.Close()
		attConn.Release()
		qd.Attachments = attachments
		questions = append(questions, qd)
	}
	chat.Questions = questions

	// answers取得
	aRows, err := conn.Query(ctx, `SELECT id, question_id, participant_id, content, created_at FROM answers WHERE chat_id = $1 ORDER BY created_at`, chatID)
	if err != nil {
		return nil, err
	}
	defer aRows.Close()
	answers := []dto.AnswerResponse{}
	for aRows.Next() {
		var ad dto.AnswerResponse
		if err := aRows.Scan(&ad.ID, &ad.QuestionID, &ad.ParticipantID, &ad.Content, &ad.CreatedAt); err != nil {
			return nil, err
		}
		// attachments取得（answerごと）
		attConn, err := q.pool.Acquire(ctx)
		if err != nil {
			return nil, err
		}
		attRows, err := attConn.Query(ctx, `SELECT type, url FROM attachments WHERE answer_id = $1`, ad.ID)
		if err != nil {
			attConn.Release()
			return nil, err
		}
		attachments := []dto.AttachmentResponse{}
		for attRows.Next() {
			var att dto.AttachmentResponse
			if err := attRows.Scan(&att.Type, &att.URL); err == nil {
				attachments = append(attachments, att)
			}
		}
		attRows.Close()
		attConn.Release()
		ad.Attachments = attachments
		answers = append(answers, ad)
	}
	chat.Answers = answers

	return &chat, nil
}

var _ interface{ query.ChatQueryInterface } = (*ChatQuery)(nil)
