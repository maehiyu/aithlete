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
	chatRepo         repository.ChatRepositoryInterface
	ParticipantRepo  repository.ParticipantRepositoryInterface
	EventPublisher   broker.ChatEventBroker
	RagRequestBroker broker.ChatEventBroker
	VectorStoreRepo  repository.VectorStoreRepositoryInterface
}

func NewChatCommandService(cr repository.ChatRepositoryInterface, pr repository.ParticipantRepositoryInterface, ep broker.ChatEventBroker, rag broker.ChatEventBroker, vsr repository.VectorStoreRepositoryInterface) *ChatCommandService {
	return &ChatCommandService{chatRepo: cr, ParticipantRepo: pr, EventPublisher: ep, RagRequestBroker: rag, VectorStoreRepo: vsr}
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

func (s *ChatCommandService) SendQuestion(chatID string, req dto.QuestionCreateRequest, token string) (*dto.QuestionResponse, error) {
	questionID := uuid.NewString()
	createdAt := time.Now()

	questionResponse := dto.QuestionResponse{
		ID:            questionID,
		ChatID:        chatID,
		ParticipantID: req.ParticipantID,
		Content:       req.Content,
		CreatedAt:     createdAt,
		Attachments:   nil,
	}

	participants, err := s.getOtherParticipants(chatID, req.ParticipantID)
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
				_ = s.RagRequestBroker.PublishChatEvent(context.Background(), event)
				break
			}
		}
	}

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
	answerID := uuid.NewString()
	createdAt := time.Now()

	answerResponse := dto.AnswerResponse{
		ID:            answerID,
		ChatID:        chatID,
		QuestionID:    req.QuestionID,
		ParticipantID: req.ParticipantID,
		Content:       req.Content,
		CreatedAt:     createdAt,
		Attachments:   nil, // 必要ならここでセット
	}

	questionContent, err := s.chatRepo.GetQuestionContent(req.QuestionID)
	if err != nil {
		return nil, err
	}

	if s.VectorStoreRepo != nil {
		_ = s.VectorStoreRepo.SaveQAPair(chatID, questionContent, answerResponse.Content, answerID)
	}

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

func (s *ChatCommandService) getOtherParticipants(chatID, excludeID string) ([]*dto.ParticipantResponse, error) {
	ids, err := s.chatRepo.GetParticipantIDsByChatID(chatID)
	if err != nil {
		return nil, err
	}
	participants := make([]*dto.ParticipantResponse, 0, len(ids))
	for _, id := range ids {
		if id == excludeID {
			continue
		}
		p, err := s.ParticipantRepo.FindByID(id)
		if err != nil || p == nil {
			continue
		}
		resp := dto.ParticipantEntityToResponse(p)
		participants = append(participants, &resp)
	}
	return participants, nil
}

func (s *ChatCommandService) SaveQuestion(q dto.QuestionResponse) error {
	questionEntity := dto.QuestionResponseToEntity(q)
	return s.chatRepo.AddQuestion(q.ChatID, questionEntity)
}

func (s *ChatCommandService) SaveAnswer(a dto.AnswerResponse) error {
	answerEntity := dto.AnswerResponseToEntity(a)
	return s.chatRepo.AddAnswer(a.ChatID, answerEntity)
}
