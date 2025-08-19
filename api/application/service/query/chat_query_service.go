package query

import (
	"api/application/dto"
	"api/application/query"
	"sort"
)

type ChatQueryService struct {
	queryService query.ChatQueryServiceInterface
}

func NewChatQueryService(qs query.ChatQueryServiceInterface) *ChatQueryService {
	return &ChatQueryService{queryService: qs}
}

func (s *ChatQueryService) FindChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error) {
	chats, err := s.queryService.FindChatsByUserID(userID)

	if err != nil {
		return nil, err
	}
	
	SortChatsByLastActive(chats)
	return chats, nil
}

func (s *ChatQueryService) FindChatByID(chatID string) (*dto.ChatDetailResponse, error) {
	return s.queryService.FindChatByID(chatID)
}

func SortChatsByLastActive(chats []dto.ChatSummaryResponse) {
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].LastActiveAt.After(chats[j].LastActiveAt)
	})
}
