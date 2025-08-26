// chats テーブル定義
// id TEXT PRIMARY KEY
// started_at TIMESTAMP NOT NULL
// last_active_at TIMESTAMP NOT NULL
// title TEXT
// participant_ids TEXT[] NOT NULL

// questions テーブル定義
// id TEXT PRIMARY KEY
// chat_id TEXT NOT NULL REFERENCES chats(id) ON DELETE CASCADE
// participant_id TEXT NOT NULL REFERENCES participants(id) ON DELETE CASCADE
// content TEXT NOT NULL
// created_at TIMESTAMP NOT NULL

// answers テーブル定義
// id TEXT PRIMARY KEY
// chat_id TEXT NOT NULL REFERENCES chats(id) ON DELETE CASCADE
// question_id TEXT NOT NULL REFERENCES questions(id) ON DELETE CASCADE
// participant_id TEXT NOT NULL REFERENCES participants(id) ON DELETE CASCADE
// content TEXT NOT NULL
// created_at TIMESTAMP NOT NULL

// attachments テーブル定義
// id TEXT PRIMARY KEY
// type TEXT NOT NULL
// url TEXT NOT NULL
// thumbnail TEXT
// pose_id TEXT
// meta TEXT
// original_id TEXT
// question_id TEXT REFERENCES questions(id) ON DELETE CASCADE
// answer_id TEXT REFERENCES answers(id) ON DELETE CASCADE

// posedata テーブル定義
// id TEXT PRIMARY KEY
// participant_ids TEXT[] NOT NULL
// score DOUBLE PRECISION
package repository

import (
	"api/domain/entity"
	"api/domain/repository"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatRepositoryImpl struct {
	conn *pgxpool.Pool
}

func NewChatRepository(conn *pgxpool.Pool) *ChatRepositoryImpl {
	return &ChatRepositoryImpl{conn: conn}
}

func (r *ChatRepositoryImpl) CreateChat(chat *entity.Chat) (*entity.Chat, error) {
	ctx := context.Background()
	conn, err := r.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// chatsテーブルInsert
	_, err = tx.Exec(ctx,
		`INSERT INTO chats (id, started_at, last_active_at, title, participant_ids) VALUES ($1, $2, $3, $4, $5)`,
		chat.ID, chat.StartedAt, chat.LastActiveAt, chat.Title, chat.ParticipantIDs,
	)
	if err != nil {
		return nil, err
	}
	// Questions, Answers, Attachments, PoseDataもInsert
	for _, q := range chat.Questions {
		_, err = tx.Exec(ctx,
			`INSERT INTO questions (id, chat_id, participant_id, content, created_at) VALUES ($1, $2, $3, $4, $5)`,
			q.ID, q.ChatID, q.ParticipantID, q.Content, q.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		for _, a := range q.Attachments {
			_, err = tx.Exec(ctx,
				`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
				a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	for _, ans := range chat.Answers {
		_, err = tx.Exec(ctx,
			`INSERT INTO answers (id, chat_id, question_id, participant_id, content, created_at) VALUES ($1,$2,$3,$4,$5,$6)`,
			ans.ID, ans.ChatID, ans.QuestionID, ans.ParticipantID, ans.Content, ans.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		for _, a := range ans.Attachments {
			_, err = tx.Exec(ctx,
				`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
				a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepositoryImpl) FindChatByID(chatId string) (*entity.Chat, error) {
	ctx := context.Background()
	var chat entity.Chat
	// chatsテーブル取得
	row := r.conn.QueryRow(ctx, `SELECT id, started_at, last_active_at, title, participant_ids FROM chats WHERE id = $1`, chatId)
	err := row.Scan(&chat.ID, &chat.StartedAt, &chat.LastActiveAt, &chat.Title, &chat.ParticipantIDs)
	if err != nil {
		return nil, err
	}

	// questions取得
	questionsRows, err := r.conn.Query(ctx, `SELECT id, chat_id, participant_id, content, created_at FROM questions WHERE chat_id = $1`, chatId)
	if err != nil {
		return nil, err
	}
	defer questionsRows.Close()
	var questions []entity.Question
	for questionsRows.Next() {
		var q entity.Question
		err := questionsRows.Scan(&q.ID, &q.ChatID, &q.ParticipantID, &q.Content, &q.CreatedAt)
		if err != nil {
			return nil, err
		}
		// attachments取得（questionごと）
		attachRows, err := r.conn.Query(ctx, `SELECT id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id FROM attachments WHERE question_id = $1`, q.ID)
		if err != nil {
			return nil, err
		}
		var attachments []entity.Attachment
		for attachRows.Next() {
			var a entity.Attachment
			err := attachRows.Scan(&a.ID, &a.Type, &a.URL, &a.Thumbnail, &a.PoseID, &a.Meta, &a.OriginalID, &a.QuestionID, &a.AnswerID)
			if err != nil {
				attachRows.Close()
				return nil, err
			}
			attachments = append(attachments, a)
		}
		attachRows.Close()
		q.Attachments = attachments
		questions = append(questions, q)
	}
	chat.Questions = questions

	// answers取得
	answersRows, err := r.conn.Query(ctx, `SELECT id, chat_id, question_id, participant_id, content, created_at FROM answers WHERE chat_id = $1`, chatId)
	if err != nil {
		return nil, err
	}
	defer answersRows.Close()
	var answers []entity.Answer
	for answersRows.Next() {
		var a entity.Answer
		err := answersRows.Scan(&a.ID, &a.ChatID, &a.QuestionID, &a.ParticipantID, &a.Content, &a.CreatedAt)
		if err != nil {
			return nil, err
		}
		// attachments取得（answerごと）
		attachRows, err := r.conn.Query(ctx, `SELECT id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id FROM attachments WHERE answer_id = $1`, a.ID)
		if err != nil {
			return nil, err
		}
		var attachments []entity.Attachment
		for attachRows.Next() {
			var att entity.Attachment
			err := attachRows.Scan(&att.ID, &att.Type, &att.URL, &att.Thumbnail, &att.PoseID, &att.Meta, &att.OriginalID, &att.QuestionID, &att.AnswerID)
			if err != nil {
				attachRows.Close()
				return nil, err
			}
			attachments = append(attachments, att)
		}
		attachRows.Close()
		a.Attachments = attachments
		answers = append(answers, a)
	}
	chat.Answers = answers

	// posedata取得（chat単位で紐付く場合のみ。なければスキップ）
	// 例: chat_idカラムがあればWHERE chat_id = $1で取得
	// chat.PoseData = ...

	return &chat, nil
}

func (r *ChatRepositoryImpl) UpdateChat(chat *entity.Chat) (*entity.Chat, error) {
	ctx := context.Background()
	conn, err := r.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	_, err = tx.Exec(ctx,
		`UPDATE chats SET started_at = $2, last_active_at = $3, title = $4, participant_ids = $5 WHERE id = $1`,
		chat.ID, chat.StartedAt, chat.LastActiveAt, chat.Title, chat.ParticipantIDs,
	)
	if err != nil {
		return nil, err
	}
	// Questions, Answers, Attachments, PoseDataもUpdate/Upsert
	for _, q := range chat.Questions {
		_, err = tx.Exec(ctx,
			`INSERT INTO questions (id, chat_id, participant_id, content, created_at) VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (id) DO UPDATE SET chat_id=EXCLUDED.chat_id, participant_id=EXCLUDED.participant_id, content=EXCLUDED.content, created_at=EXCLUDED.created_at`,
			q.ID, q.ChatID, q.ParticipantID, q.Content, q.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		for _, a := range q.Attachments {
			_, err = tx.Exec(ctx,
				`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
				ON CONFLICT (id) DO UPDATE SET type=EXCLUDED.type, url=EXCLUDED.url, thumbnail=EXCLUDED.thumbnail, pose_id=EXCLUDED.pose_id, meta=EXCLUDED.meta, original_id=EXCLUDED.original_id, question_id=EXCLUDED.question_id, answer_id=EXCLUDED.answer_id`,
				a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	for _, ans := range chat.Answers {
		_, err = tx.Exec(ctx,
			`INSERT INTO answers (id, chat_id, question_id, participant_id, content, created_at) VALUES ($1,$2,$3,$4,$5,$6)
			ON CONFLICT (id) DO UPDATE SET chat_id=EXCLUDED.chat_id, question_id=EXCLUDED.question_id, participant_id=EXCLUDED.participant_id, content=EXCLUDED.content, created_at=EXCLUDED.created_at`,
			ans.ID, ans.ChatID, ans.QuestionID, ans.ParticipantID, ans.Content, ans.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		for _, a := range ans.Attachments {
			_, err = tx.Exec(ctx,
				`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
				ON CONFLICT (id) DO UPDATE SET type=EXCLUDED.type, url=EXCLUDED.url, thumbnail=EXCLUDED.thumbnail, pose_id=EXCLUDED.pose_id, meta=EXCLUDED.meta, original_id=EXCLUDED.original_id, question_id=EXCLUDED.question_id, answer_id=EXCLUDED.answer_id`,
				a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
			)
			if err != nil {
				return nil, err
			}
		}
	}
	// PoseDataはAttachment経由で管理するため、ここではUpsertしない
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	return chat, nil
}

func (r *ChatRepositoryImpl) AddQuestion(chatId string, question *entity.Question) (*entity.Chat, error) {
	ctx := context.Background()
	conn, err := r.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Insert question
	_, err = tx.Exec(ctx,
		`INSERT INTO questions (id, chat_id, participant_id, content, created_at) VALUES ($1, $2, $3, $4, $5)`,
		question.ID, chatId, question.ParticipantID, question.Content, question.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	// Insert attachments (if any)
	for _, a := range question.Attachments {
		_, err = tx.Exec(ctx,
			`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
			a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
		)
		if err != nil {
			return nil, err
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	// Return updated chat
	return r.FindChatByID(chatId)
}

func (r *ChatRepositoryImpl) AddAnswer(chatId string, answer *entity.Answer) (*entity.Chat, error) {
	ctx := context.Background()
	conn, err := r.conn.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		}
	}()

	// Insert answer
	_, err = tx.Exec(ctx,
		`INSERT INTO answers (id, chat_id, question_id, participant_id, content, created_at) VALUES ($1,$2,$3,$4,$5,$6)`,
		answer.ID, chatId, answer.QuestionID, answer.ParticipantID, answer.Content, answer.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	// Insert attachments (if any)
	for _, a := range answer.Attachments {
		_, err = tx.Exec(ctx,
			`INSERT INTO attachments (id, type, url, thumbnail, pose_id, meta, original_id, question_id, answer_id) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
			a.ID, a.Type, a.URL, a.Thumbnail, a.PoseID, a.Meta, a.OriginalID, a.QuestionID, a.AnswerID,
		)
		if err != nil {
			return nil, err
		}
	}
	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}
	// Return updated chat
	return r.FindChatByID(chatId)
}

// domain/repository.ChatRepositoryインターフェースを満たす
var _ repository.ChatRepositoryInterface = (*ChatRepositoryImpl)(nil)
