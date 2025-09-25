package repository_test

import (
	"api/domain/entity"
	"api/infrastructure/repository"
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestDB initializes a connection pool for the test database.
func setupTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()
	dsn := os.Getenv("TEST_DB_DSN")
	if dsn == "" {
		t.Skip("TEST_DB_DSN environment variable is not set, skipping integration tests.")
	}
	pool, err := pgxpool.New(context.Background(), dsn)
	require.NoError(t, err, "failed to connect to test db")
	return pool
}

// cleanup truncates the participants table to ensure a clean state.
func cleanup(t *testing.T, pool *pgxpool.Pool) {
	t.Helper()
	_, err := pool.Exec(context.Background(), "TRUNCATE TABLE participants RESTART IDENTITY")
	require.NoError(t, err, "failed to clean participants table")
}

func TestParticipantRepository(t *testing.T) {
	pool := setupTestDB(t)
	defer pool.Close()
	repo := repository.NewParticipantRepository(pool)

	// Ensure a clean state before all tests in this function.
	cleanup(t, pool)

	t.Run("Create and FindByID", func(t *testing.T) {
		iconURL := "http://example.com/icon.png"
		p := &entity.Participant{
			ID:      "user-1",
			Name:    "Test User 1",
			Email:   "test1@example.com",
			Role:    "user",
			Sports:  []string{"soccer"},
			IconURL: &iconURL,
		}

		// Act: Create
		createdID, err := repo.Create(p)
		require.NoError(t, err)
		assert.Equal(t, p.ID, createdID)

		// Act: Find
		found, err := repo.FindByID(p.ID)

		// Assert
		require.NoError(t, err)
		assert.Equal(t, p.ID, found.ID)
		assert.Equal(t, p.Name, found.Name)
		assert.Equal(t, p.Email, found.Email)
		assert.Equal(t, p.Role, found.Role)
		assert.Equal(t, p.Sports, found.Sports)
		assert.Equal(t, p.IconURL, found.IconURL)
	})

	t.Run("FindByID - Not Found", func(t *testing.T) {
		// Act
		_, err := repo.FindByID("non-existent-id")

		// Assert
		assert.Error(t, err)
	})

	t.Run("Update", func(t *testing.T) {
		// Arrange: Create a participant to update
		p := &entity.Participant{ID: "user-2", Name: "Before Update", Email: "update@example.com", Role: "user"}
		_, err := repo.Create(p)
		require.NoError(t, err)

		// Act
		updatedName := "After Update"
		p.Name = updatedName
		err = repo.Update(p)
		require.NoError(t, err)

		// Assert
		found, err := repo.FindByID(p.ID)
		require.NoError(t, err)
		assert.Equal(t, updatedName, found.Name)
	})

	t.Run("FindManyByIDs", func(t *testing.T) {
		// Arrange: Create multiple participants
		p1 := &entity.Participant{ID: "user-3", Name: "User 3", Email: "3@example.com", Role: "user"}
		p2 := &entity.Participant{ID: "user-4", Name: "User 4", Email: "4@example.com", Role: "user"}
		_, err := repo.Create(p1)
		require.NoError(t, err)
		_, err = repo.Create(p2)
		require.NoError(t, err)

		// Act
		idsToFind := []string{p1.ID, p2.ID, "non-existent"}
		found, err := repo.FindByIDs(idsToFind)

		// Assert
		require.NoError(t, err)
		assert.Len(t, found, 2) // Should only find the two created participants

		// Check that the correct participants were returned
		foundIDs := make([]string, len(found))
		for i, p := range found {
			foundIDs[i] = p.ID
		}
		assert.ElementsMatch(t, []string{p1.ID, p2.ID}, foundIDs)
	})
}