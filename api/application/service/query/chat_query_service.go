package query

import (
	"api/application/dto"
	"api/application/query"
)

type ChatQueryService struct {
	chatQuery        query.ChatQueryInterface
	participantQuery query.ParticipantQueryInterface
}

func NewChatQueryService(chatQuery query.ChatQueryInterface, participantQuery query.ParticipantQueryInterface) *ChatQueryService {
	return &ChatQueryService{chatQuery: chatQuery, participantQuery: participantQuery}
}
	
func (s *ChatQueryService) GetChatsByUserID(userID string) ([]dto.ChatSummaryResponse, error) {
	chats, err := s.chatQuery.FindChatsByUserID(userID)

	if err != nil {
		return nil, err
	}
	
	return chats, nil
}

func (s *ChatQueryService) GetChatByID(chatID string) (*dto.ChatDetailResponse, error) {
	chat, err := s.chatQuery.FindChatByID(chatID)
	if err != nil {
		return nil, err
	}

	participants, err := s.participantQuery.FindParticipantsByIDs(chat.ParticipantIDs)
	if err != nil {
		return nil, err
	}

	response := dto.ChatEntityToDetailResponse(chat, participants)
	return response, nil
}