package rag

type RagClient interface {
	CallRAGServer(chatID, questionContent, aiID, questionID, token string) (string, error)
}