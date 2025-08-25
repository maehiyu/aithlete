package repository_test

import (
	"context"
	"os"
	"testing"

	"api/domain/entity"
	"api/infrastructure/repository"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

func setupTestParticipantPool(t *testing.T) *pgxpool.Pool {
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

func TestParticipantRepositoryImpl_CRUD(t *testing.T) {
	pool := setupTestParticipantPool(t)
	defer pool.Close()
	repo := repository.NewParticipantRepository(pool)

	p := &entity.Participant{
		ID:     "test_p1",
		Name:   "Test User",
		Email:  "test@example.com",
		Role:   "player",
		Sports: []string{"soccer", "tennis"},
	}
	created, err := repo.Create(p)
	assert.NoError(t, err)
	assert.Equal(t, "test_p1", created.ID)

	found, err := repo.FindByID("test_p1")
	assert.NoError(t, err)
	assert.Equal(t, "Test User", found.Name)
	assert.ElementsMatch(t, []string{"soccer", "tennis"}, found.Sports)

	found.Name = "Updated User"
	found.Sports = []string{"baseball"}
	updated, err := repo.Update(found)
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", updated.Name)
	assert.ElementsMatch(t, []string{"baseball"}, updated.Sports)

	found2, err := repo.FindByID("test_p1")
	assert.NoError(t, err)
	assert.Equal(t, "Updated User", found2.Name)
	assert.ElementsMatch(t, []string{"baseball"}, found2.Sports)
}
