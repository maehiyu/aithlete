package repository

type VectorStoreRepositoryInterface interface {
	SaveQAPair(chatID, question, answer, answerID string) error
}