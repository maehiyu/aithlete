package repository_test

import (
	repo "api/infrastructure/repository"
	"context"
	"os"
	"testing"

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
	repository := repo.NewChatRepository(pool)

	// 既定データを利用
	chatID := "chat1"
	questionID := "q1"
	answerID := "a1"

	found, err := repository.FindChatByID(chatID)
	if err != nil {
		t.Fatalf("FindChatByID failed: %v", err)
	}
	if found.ID != chatID {
		t.Errorf("expected ID %s, got %s", chatID, found.ID)
	}

	t.Run("CheckQuestionExists", func(t *testing.T) {
		foundQ := false
		for _, qq := range found.Questions {
			if qq.ID == questionID {
				foundQ = true
				break
			}
		}
		if !foundQ {
			t.Errorf("question %s not found in chat %s", questionID, chatID)
		}
	})

	t.Run("CheckAnswerExists", func(t *testing.T) {
		foundA := false
		for _, aa := range found.Answers {
			if aa.ID == answerID {
				foundA = true
				break
			}
		}
		if !foundA {
			t.Errorf("answer %s not found in chat %s", answerID, chatID)
		}
	})
}

func TestChatRepository_UpdateChat(t *testing.T) {
	pool := setupTestRepositoryPool(t)
	defer pool.Close()
	repository := repo.NewChatRepository(pool)

	chatID := "chat1"
	found, err := repository.FindChatByID(chatID)
	if err != nil {
		t.Fatalf("FindChatByID failed: %v", err)
	}

	title := "updated title"
	found.Title = &title
	err = repository.UpdateChat(found)
	if err != nil {
		t.Fatalf("UpdateChat failed: %v", err)
	}

	origTitle := "テストチャット"
	found.Title = &origTitle
	_ = repository.UpdateChat(found)
}
