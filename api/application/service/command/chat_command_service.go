package command

import (
	"api/application/dto"
	"api/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ChatCommandService struct {
	chatRepo                  repository.ChatRepositoryInterface
	ParticipantRepo            repository.ParticipantRepositoryInterface
}

func NewChatCommandService(cr repository.ChatRepositoryInterface, pr repository.ParticipantRepositoryInterface) *ChatCommandService {
	return &ChatCommandService{chatRepo: cr, ParticipantRepo: pr}
}

func (s *ChatCommandService) CreateChat(chat dto.ChatCreateRequest) (*dto.ChatDetailResponse, error) {
	chatEntity := dto.ChatCreateRequestToEntity(chat, uuid.NewString(), time.Now())

	savedChat, err := s.chatRepo.CreateChat(chatEntity)
	if err != nil {
		return nil, err
	}

	participants := make([]dto.ParticipantResponse, 0, len(chat.ParticipantIDs))

	for _, id := range chat.ParticipantIDs {
		p, err := s.ParticipantRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		participants = append(participants, dto.ParticipantEntityToResponse(p))
	}

	chatDetail := dto.ChatEntityToDetailResponse(savedChat, participants)
	return &chatDetail, nil
}


func (s *ChatCommandService) UpdateChat(chatID string, chat dto.ChatUpdateRequest) (*dto.ChatDetailResponse, error) {
	chatEntity, err := s.chatRepo.FindChatByID(chatID)
	if err != nil {
		return nil, err
	}
	
	// DTOの値を適用
	dto.ChatUpdateRequestToEntity(chatEntity, chat)

	// 保存
	updatedChat, err := s.chatRepo.UpdateChat(chatEntity)
	if err != nil {
		return nil, err
	}
	
	participants := make([]dto.ParticipantResponse, 0, len(updatedChat.ParticipantIDs))

	for _, id := range updatedChat.ParticipantIDs {
		p, err := s.ParticipantRepo.FindByID(id)
		if err != nil {
			return nil, err
		}
		participants = append(participants, dto.ParticipantEntityToResponse(p))
	}
	
	chatDetail := dto.ChatEntityToDetailResponse(updatedChat, participants)
	return &chatDetail, nil
}



