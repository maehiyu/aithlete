//go:generate mockgen -source=vector_store_repository.go -destination=mocks/mock_vector_store_repository.go -package=mocks VectorStoreRepositoryInterface
package repository

type VectorStoreRepositoryInterface interface {
	SaveQAPair(chatID, question, answer, answerID string) error
}