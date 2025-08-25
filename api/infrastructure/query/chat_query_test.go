package query_test

import (
	"context"
	"os"
	"testing"

	infraquery "api/infrastructure/query"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func setupTestQueryPool(t *testing.T) *pgxpool.Pool {
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		t.Fatal("TEST_DB_DSN environment variable is not set")
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	return pool
}

func TestChatQuery_FindChatsByUserID(t *testing.T) {
	pool := setupTestQueryPool(t)
	defer pool.Close()
	q := infraquery.NewChatQuery(pool)

	// user1, chat1 などinit.sqlのテストデータを利用
	chats, err := q.FindChatsByUserID("user1")
	assert.NoError(t, err)
	assert.NotEmpty(t, chats)
	assert.Equal(t, "chat1", chats[0].ID)
}

func TestChatQuery_FindChatByID(t *testing.T) {
	pool := setupTestQueryPool(t)
	defer pool.Close()
	q := infraquery.NewChatQuery(pool)

	chat, err := q.FindChatByID("chat1")
	assert.NoError(t, err)
	assert.NotNil(t, chat)
	assert.Equal(t, "chat1", chat.ID)
	assert.NotEmpty(t, chat.Participants)
	assert.NotEmpty(t, chat.Questions)
	assert.NotEmpty(t, chat.Answers)
}
