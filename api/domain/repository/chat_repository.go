//go:generate mockgen -source=chat_repository.go -destination=mocks/mock_chat_repository.go -package=mocks ChatRepositoryInterface
package repository

import (
	"api/domain/entity"
	"context"
)

type ChatRepositoryInterface interface {
	CreateChat(ctx context.Context, chat *entity.Chat) (string, error)
	FindChatByID(ctx context.Context, chatId string) (*entity.Chat, error)
	
	UpdateChat(ctx context.Context, chat *entity.Chat) error
	AddQuestion(ctx context.Context, chatId string, question *entity.Question) error
	AddAnswer(ctx context.Context, chatId string, answer *entity.Answer) error
	FindParticipantIDsByChatID(ctx context.Context, chatId string) ([]string, error)
	GetQuestionContent(ctx context.Context, questionID string) (string, error)
}