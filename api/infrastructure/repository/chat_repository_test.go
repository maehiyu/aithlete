package repository_test

import (
	"context"
	"os"
	"testing"
	"time"

	"api/domain/entity"
	repo "api/infrastructure/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func setupTestRepositoryPool(t *testing.T) *pgxpool.Pool {
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

func TestChatRepository_CreateAndFind(t *testing.T) {
	pool := setupTestRepositoryPool(t)
	defer pool.Close()
	// PoolではトランザクションはConn単位で取得する必要があるが、
	// ここでは簡易的にPoolを直接使う
	repository := repo.NewChatRepository(pool)

	chat := &entity.Chat{
		ID:             "test_chat1",
		StartedAt:      time.Now(),
		LastActiveAt:   time.Now(),
		Title:          nil,
		ParticipantIDs: []string{"test_user1", "test_user2"},
	}

	created, err := repository.CreateChat(chat)
	if err != nil {
		t.Fatalf("CreateChat failed: %v", err)
	}
	if created.ID != chat.ID {
		t.Errorf("expected ID %s, got %s", chat.ID, created.ID)
	}

	found, err := repository.FindChatByID("test_chat1")
	if err != nil {
		t.Fatalf("FindChatByID failed: %v", err)
	}
	if found.ID != chat.ID {
		t.Errorf("expected ID %s, got %s", chat.ID, found.ID)
	}
}

func TestChatRepository_UpdateChat(t *testing.T) {
	pool := setupTestRepositoryPool(t)
	defer pool.Close()
	repository := repo.NewChatRepository(pool)

	chat := &entity.Chat{
		ID:             "test_chat2",
		StartedAt:      time.Now(),
		LastActiveAt:   time.Now(),
		Title:          nil,
		ParticipantIDs: []string{"test_user1", "test_user2"},
	}
	if _, err := repository.CreateChat(chat); err != nil {
		t.Fatalf("CreateChat failed: %v", err)
	}

	title := "updated title"
	chat.Title = &title
	updated, err := repository.UpdateChat(chat)
	if err != nil {
		t.Fatalf("UpdateChat failed: %v", err)
	}
	if updated.Title == nil || *updated.Title != title {
		t.Errorf("expected title %s, got %v", title, updated.Title)
	}
}
