//go:generate mockgen -source=vector_store_repository.go -destination=mocks/mock_vector_store_repository.go -package=mocks VectorStoreRepositoryInterface
package repository

import (
	"context"
)

type VectorStoreRepositoryInterface interface {
	SaveQAPair(ctx context.Context, chatID, question, answer, answerID string) error
}