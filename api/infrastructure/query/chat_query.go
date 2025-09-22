package query

import (
	"api/application/dto" 
	"api/application/query"
	"api/domain/entity" 
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatQuery struct {
	pool *pgxpool.Pool
}

func NewChatQuery(pool *pgxpool.Pool) *ChatQuery {
	return &ChatQuery{pool: pool}
}

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


func (q *ChatQuery) FindChatByID(chatID string) (*entity.Chat, error) {
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

	var chat entity.Chat
	var title *string
	var participantIDs []string
	if err := row.Scan(&chat.ID, &title, &chat.StartedAt, &chat.LastActiveAt, &participantIDs); err != nil {
		return nil, err
	}
	chat.Title = title
	chat.ParticipantIDs = participantIDs

	qRows, err := conn.Query(ctx, `SELECT id, chat_id, participant_id, content, created_at FROM questions WHERE chat_id = $1 ORDER BY created_at`, chatID)
	if err != nil {
		return nil, err
	}
	defer qRows.Close()
	questions := []entity.Question{}
	for qRows.Next() {
		var qItem entity.Question
		if err := qRows.Scan(&qItem.ID, &qItem.ChatID, &qItem.ParticipantID, &qItem.Content, &qItem.CreatedAt); err != nil {
			return nil, err
		}

		var attConn *pgxpool.Conn 
		attConn, err = q.pool.Acquire(ctx)
		if err != nil {
			return nil, err
		}
	
		aRows, err := attConn.Query(ctx, `SELECT id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id FROM attachments WHERE question_id = $1`, qItem.ID)
		if err != nil {
			attConn.Release() 
			return nil, err
		}
		attachments := []entity.Attachment{} 
		for aRows.Next() {
			var a entity.Attachment
			if err := aRows.Scan(&a.ID, &a.Type, &a.URL, &a.Thumbnail, &a.PoseID, &a.Meta, &a.OriginalID, &a.QuestionID, &a.AnswerID); err != nil {
				aRows.Close()
				attConn.Release()
				return nil, err
			}
			attachments = append(attachments, a)
		}
		aRows.Close()
		attConn.Release()
		qItem.Attachments = attachments
		questions = append(questions, qItem)
	}
	chat.Questions = questions

	aRows, err := conn.Query(ctx, `SELECT id, chat_id, question_id, participant_id, content, created_at FROM answers WHERE chat_id = $1 ORDER BY created_at`, chatID)
	if err != nil {
		return nil, err
	}
	defer aRows.Close()
	answers := []entity.Answer{} 
	for aRows.Next() {
		var aItem entity.Answer 
		if err := aRows.Scan(&aItem.ID, &aItem.ChatID, &aItem.QuestionID, &aItem.ParticipantID, &aItem.Content, &aItem.CreatedAt); err != nil {
			return nil, err
		}

		var attConn *pgxpool.Conn
		attConn, err = q.pool.Acquire(ctx)
		if err != nil {
			return nil, err
		}
	
		attRows, err := attConn.Query(ctx, `SELECT id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id FROM attachments WHERE answer_id = $1`, aItem.ID) 
		if err != nil {
			attConn.Release() 
			return nil, err
		}
		attachments := []entity.Attachment{}
		for attRows.Next() {
			var att entity.Attachment
			if err := attRows.Scan(&att.ID, &att.Type, &att.URL, &att.Thumbnail, &att.PoseID, &att.Meta, &att.OriginalID, &att.QuestionID, &att.AnswerID); err != nil {
				attRows.Close()
				attConn.Release()
				return nil, err
			}
			attachments = append(attachments, att)
		}
		attRows.Close()
		attConn.Release()
		aItem.Attachments = attachments
		answers = append(answers, aItem)
	}
	chat.Answers = answers

	return &chat, nil
}

var _ query.ChatQueryInterface = (*ChatQuery)(nil)
