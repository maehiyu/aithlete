//go:generate mockgen -source=chat_repository.go -destination=mocks/mock_chat_repository.go -package=mocks ChatRepositoryInterface
package repository

import (
	"api/domain/entity"
)

type ChatRepositoryInterface interface {
	CreateChat(chat *entity.Chat) (string, error)
	FindChatByID(chatId string) (*entity.Chat, error)
	
	UpdateChat(chat *entity.Chat) error
	AddQuestion(chatId string, question *entity.Question) error
	AddAnswer(chatId string, answer *entity.Answer) error
	FindParticipantIDsByChatID(chatId string) ([]string, error)
	GetQuestionContent(questionID string) (string, error)
}