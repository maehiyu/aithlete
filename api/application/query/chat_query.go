//go:generate mockgen -source=chat_query.go -destination=mocks/mock_chat_query.go -package=mocks ChatQueryInterface
package query

import (
	"api/application/dto"
	"api/domain/entity"
)

type ChatQueryInterface interface {
	FindChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error)
	FindChatByID(chatID string) (*entity.Chat, error)
}