package query

import "api/application/dto"

type ChatQueryInterface interface {
	FindChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error)
	FindChatByID(chatID string) (*dto.ChatDetailResponse, error)
}