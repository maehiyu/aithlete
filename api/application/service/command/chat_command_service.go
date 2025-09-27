package command

import (
	"api/application/broker"
	"api/application/dto"
	"api/domain/entity"
	"api/domain/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type ChatCommandService struct {
	chatRepo         repository.ChatRepositoryInterface
	ParticipantRepo  repository.ParticipantRepositoryInterface
	EventPublisher   broker.ChatEventBroker
	RagRequestBroker broker.ChatEventBroker
	VectorStoreRepo  repository.VectorStoreRepositoryInterface
}

func NewChatCommandService(cr repository.ChatRepositoryInterface, pr repository.ParticipantRepositoryInterface, ep broker.ChatEventBroker, rag broker.ChatEventBroker, vsr repository.VectorStoreRepositoryInterface) *ChatCommandService {
	return &ChatCommandService{chatRepo: cr, ParticipantRepo: pr, EventPublisher: ep, RagRequestBroker: rag, VectorStoreRepo: vsr}
}

func (s *ChatCommandService) CreateChat(ctx context.Context, chat dto.ChatCreateRequest, userID string) (string, error) {
	chat.ParticipantIDs = append(chat.ParticipantIDs, userID)
	chatEntity := dto.ChatCreateRequestToEntity(chat, uuid.NewString(), time.Now())

	chatID, err := s.chatRepo.CreateChat(ctx, chatEntity)
	if err != nil {
		return "", err
	}

	return chatID, nil
}

func (s *ChatCommandService) UpdateChat(ctx context.Context, chatID string, chat dto.ChatUpdateRequest) error {
	chatEntity, err := s.chatRepo.FindChatByID(ctx, chatID)
	if err != nil {
		return err
	}

	dto.ChatUpdateRequestToEntity(chatEntity, chat)

	err = s.chatRepo.UpdateChat(ctx, chatEntity)
	if err != nil {
		return err
	}
	return nil
}

func (s *ChatCommandService) SendMessage(ctx context.Context, chatID string, req dto.ChatItemRequest) error {
	id := uuid.NewString()
	createdAt := time.Now()

	questionResponse := dto.ChatItem{
		ID:            id,
		ChatID:        chatID,
		ParticipantID: req.ParticipantID,
		Content:       req.Content,
		CreatedAt:     createdAt,
		Type:          req.Type,
		QuestionID:    req.QuestionID,
		Attachments:   nil,
		TempID:        &req.TempID,
	}

	// Publish event to all participants (including self for saving)
	allParticipantIDs, err := s.chatRepo.FindParticipantIDsByChatID(ctx, chatID)
	if err == nil && s.EventPublisher != nil {
		event := dto.ChatEvent{
			ID:        uuid.NewString(),
			ChatID:    chatID,
			Type:      "message",
			From:      req.ParticipantID,
			To:        allParticipantIDs,
			Timestamp: time.Now().Unix(),
			Payload:   questionResponse,
		}
		_ = s.EventPublisher.PublishChatEvent(ctx, event)
	}

	// Publish event to RAG service if ai_coach exists
	participants, err := s.getOtherParticipants(ctx, chatID, req.ParticipantID)
	if err == nil && s.RagRequestBroker != nil {
		for _, p := range participants {
			if p.Role == "ai_coach" {
				event := dto.ChatEvent{
					ID:        uuid.NewString(),
					ChatID:    chatID,
					Type:      "rag_request",
					From:      req.ParticipantID,
					To:        []string{p.ID},
					Timestamp: time.Now().Unix(),
					Payload:   questionResponse,
				}
				_ = s.RagRequestBroker.PublishChatEvent(ctx, event)
				break
			}
		}
	}


	return nil
}

func (s *ChatCommandService) getOtherParticipantIDs(ctx context.Context, chatID, excludeID string) ([]string, error) {
	ids, err := s.chatRepo.FindParticipantIDsByChatID(ctx, chatID)
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

func (s *ChatCommandService) getOtherParticipants(ctx context.Context, chatID, excludeID string) ([]*entity.Participant, error) {
	ids, err := s.chatRepo.FindParticipantIDsByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	participants, err := s.ParticipantRepo.FindByIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (s *ChatCommandService) SaveMessage(ctx context.Context, item dto.ChatItem) error {
	switch item.Type {
	case "question":
		q := dto.ChatItemToQuestion(item)
		return s.chatRepo.AddQuestion(ctx, item.ChatID, q)
	case "answer":
		a := dto.ChatItemToAnswer(item)
		return s.chatRepo.AddAnswer(ctx, item.ChatID, a)
	case "ai_answer":
		a := dto.ChatItemToAnswer(item)
		if s.VectorStoreRepo != nil {
			questionContent, err := s.chatRepo.GetQuestionContent(ctx, a.QuestionID)
			if err == nil {
				_ = s.VectorStoreRepo.SaveQAPair(ctx, a.ChatID, questionContent, a.Content, a.ID)
			}
		}
		return s.chatRepo.AddAnswer(ctx, item.ChatID, a)
	default:
		return nil
	}
}
