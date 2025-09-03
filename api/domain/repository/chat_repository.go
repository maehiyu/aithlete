package repository

import (
	"api/domain/entity"
)

type ChatRepositoryInterface interface {
	CreateChat(chat *entity.Chat) (*entity.Chat, error)
	FindChatByID(chatId string) (*entity.Chat, error)
	UpdateChat(chat *entity.Chat) (*entity.Chat, error)
	AddQuestion(chatId string, question *entity.Question) error
	AddAnswer(chatId string, answer *entity.Answer) error
	GetParticipantIDsByChatID(chatId string) ([]string, error)
	GetQuestionContent(questionID string) (string, error)
}