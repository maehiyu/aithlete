package query

import (
	"api/application/dto"
	"api/application/query"
	"sort"
)

type ChatQueryService struct {
	chatQuery query.ChatQueryInterface
}

func NewChatQueryService(qs query.ChatQueryInterface) *ChatQueryService {
	return &ChatQueryService{chatQuery: qs}
}
	
func (s *ChatQueryService) GetChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error) {
	chats, err := s.chatQuery.FindChatsByUserID(userID)

	if err != nil {
		return nil, err
	}
	
	SortChatsByLastActive(chats)
	return chats, nil
}

func (s *ChatQueryService) GetChatByID(chatID string) (*dto.ChatDetailResponse, error) {
	return s.chatQuery.FindChatByID(chatID)
}

func SortChatsByLastActive(chats []dto.ChatSummaryResponse) {
	sort.Slice(chats, func(i, j int) bool {
		return chats[i].LastActiveAt.After(chats[j].LastActiveAt)
	})
}
