package command

import (
	"api/application/dto"
	"api/domain/entity"
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
	chatEntity := NewChatFromDTO(chat, uuid.NewString(), time.Now())

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
		participants = append(participants, *p)
	}

	chatDetail := NewChatDetailResponseFromEntity(savedChat, participants)
	return chatDetail, nil
}


func (s *ChatCommandService) UpdateChat(chat dto.ChatUpdateRequest) (*dto.ChatDetailResponse, error) {
	chatEntity, err := s.chatRepo.FindChatByID(chat.ID)
	if err != nil {
		return nil, err
	}
	
	// DTOの値を適用
	ApplyChatUpdateDTOToEntity(chatEntity, chat)
	
	// 保存
	updatedChat, err := s.chatRepo.UpdateChat(chatEntity)
	if err != nil {
		return nil, err
	}
	
	participants := make([]dto.ParticipantResponse, 0, len(updatedChat.ParticipantIDs))
	for _, id := range updatedChat.ParticipantIDs {
		p, err := s.ParticipantRepo.FindParticipantByID(id)
		if err != nil {
			return nil, err
		}
		participants = append(participants, *p)
	}
	
	chatDetail := NewChatDetailResponseFromEntity(updatedChat, participants)
	return chatDetail, nil
}

func NewChatFromDTO(dto dto.ChatCreateRequest, id string, now time.Time) *entity.Chat {
	return &entity.Chat{
		ID:            id,
		Title:         dto.Title,
		ParticipantIDs: dto.ParticipantIDs,
		Questions:     []entity.Question{},
		Answers:       []entity.Answer{},
		StartedAt:     now,
		LastActiveAt:  now,
	}
}

func NewChatDetailResponseFromEntity(chat *entity.Chat, participants []dto.ParticipantResponse) *dto.ChatDetailResponse {
	return &dto.ChatDetailResponse{
		ID:           chat.ID,
		Title:        chat.Title,
		Participants: participants,
		Questions:    []dto.QuestionDetailResponse{}, // 必要に応じて変換
		Answers:      []dto.AnswerDetailResponse{},   // 必要に応じて変換
		StartedAt:    chat.StartedAt,
		LastActiveAt: chat.LastActiveAt,
	}
}

func ApplyChatUpdateDTOToEntity(chat *entity.Chat, dto dto.ChatUpdateRequest) {
	if dto.Title != nil {
		chat.Title = dto.Title
	}
	if dto.ParticipantIDs != nil {
		chat.ParticipantIDs = dto.ParticipantIDs
	}
}