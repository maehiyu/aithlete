//go:generate mockgen -source=chat_query.go -destination=mocks/mock_chat_query.go -package=mocks ChatQueryInterface
package query

import (
	"api/application/dto"
	"api/domain/entity"
	"context"
)

type ChatQueryInterface interface {
	FindChatsByUserID(ctx context.Context, userID string) ([]dto.ChatSummaryResponse, error)
	FindChatByID(ctx context.Context, chatID string) (*entity.Chat, error)
}