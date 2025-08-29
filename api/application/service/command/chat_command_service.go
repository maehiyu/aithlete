package command

import (
	"api/application/broker"
	"api/application/dto"
	"api/domain/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type ChatCommandService struct {
	chatRepo        repository.ChatRepositoryInterface
	ParticipantRepo repository.ParticipantRepositoryInterface
	EventPublisher  broker.ChatEventPublisher
}

func NewChatCommandService(cr repository.ChatRepositoryInterface, pr repository.ParticipantRepositoryInterface, ep broker.ChatEventPublisher) *ChatCommandService {
	return &ChatCommandService{chatRepo: cr, ParticipantRepo: pr, EventPublisher: ep}
}

func (s *ChatCommandService) CreateChat(chat dto.ChatCreateRequest, userID string) (*dto.ChatDetailResponse, error) {
	chat.ParticipantIDs = append(chat.ParticipantIDs, userID)
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

func (s *ChatCommandService) SendQuestion(chatID string, req dto.QuestionCreateRequest) (*dto.QuestionResponse, error) {
	questionEntity := dto.QuestionCreateRequestToEntity(req, uuid.NewString(), chatID, time.Now())

	err := s.chatRepo.AddQuestion(chatID, questionEntity)
	if err != nil {
		return nil, err
	}

	questionResponse := dto.QuestionEntityToResponse(questionEntity)

	filtered, err := s.getOtherParticipantIDs(chatID, req.ParticipantID)
	if err != nil {
		return nil, err
	}

	if s.EventPublisher != nil {
		event := dto.ChatEvent{
			ID:        uuid.NewString(),
			ChatID:    chatID,
			Type:      "question",
			From:      req.ParticipantID,
			To:        filtered,
			Timestamp: time.Now().Unix(),
			Payload:   questionResponse,
		}
		_ = s.EventPublisher.PublishChatEvent(context.Background(), event)
	}

	return &questionResponse, nil
}

func (s *ChatCommandService) SendAnswer(chatID string, req dto.AnswerCreateRequest) (*dto.AnswerResponse, error) {
	answerEntity := dto.AnswerCreateRequestToEntity(req, uuid.NewString(), chatID, req.QuestionID, time.Now())

	err := s.chatRepo.AddAnswer(chatID, answerEntity)
	if err != nil {
		return nil, err
	}

	answerResponse := dto.AnswerEntityToResponse(answerEntity)

	filtered, err := s.getOtherParticipantIDs(chatID, req.ParticipantID)
	if err != nil {
		return nil, err
	}

	if s.EventPublisher != nil {
		event := dto.ChatEvent{
			ID:        uuid.NewString(),
			ChatID:    chatID,
			Type:      "answer",
			From:      req.ParticipantID,
			To:        filtered,
			Timestamp: time.Now().Unix(),
			Payload:   answerResponse,
		}
		_ = s.EventPublisher.PublishChatEvent(context.Background(), event)
	}

	return &answerResponse, nil
}

func (s *ChatCommandService) getOtherParticipantIDs(chatID, excludeID string) ([]string, error) {
	ids, err := s.chatRepo.GetParticipantIDsByChatID(chatID)
	if err != nil {
		return nil, err
	}
	filtered := make([]string, 0, len(ids))
	for _, id := range ids {
		if id != excludeID {
			filtered = append(filtered, id)
		}
	}
	return filtered, nil
}
