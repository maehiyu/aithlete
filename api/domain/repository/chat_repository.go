package repository

import (
	"api/domain/entity"
)

type ChatRepositoryInterface interface {
	CreateChat(chat *entity.Chat) (*entity.Chat, error)
	FindChatByID(chatId string) (*entity.Chat, error)
	UpdateChat(chat *entity.Chat) (*entity.Chat, error)
	AddQuestion(chatId string, question *entity.Question) (*entity.Chat, error)
	AddAnswer(chatId string, answer *entity.Answer) (*entity.Chat, error)
}