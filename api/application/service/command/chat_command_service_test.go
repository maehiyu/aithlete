package command

import (
	"api/application/broker/mocks"
	"api/application/dto"
	"api/domain/entity"
	domain_mocks "api/domain/repository/mocks"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// setupService is a helper function to create a new ChatCommandService with all its dependencies mocked.
func setupService(t *testing.T) (*ChatCommandService, *domain_mocks.MockChatRepositoryInterface, *domain_mocks.MockParticipantRepositoryInterface, *mocks.MockChatEventBroker, *mocks.MockChatEventBroker, *domain_mocks.MockVectorStoreRepositoryInterface) {
	ctrl := gomock.NewController(t)

	mockChatRepo := domain_mocks.NewMockChatRepositoryInterface(ctrl)
	mockParticipantRepo := domain_mocks.NewMockParticipantRepositoryInterface(ctrl)
	mockEventBroker := mocks.NewMockChatEventBroker(ctrl)
	mockRagBroker := mocks.NewMockChatEventBroker(ctrl)
	mockVectorRepo := domain_mocks.NewMockVectorStoreRepositoryInterface(ctrl)

	service := NewChatCommandService(mockChatRepo, mockParticipantRepo, mockEventBroker, mockRagBroker, mockVectorRepo)

	return service, mockChatRepo, mockParticipantRepo, mockEventBroker, mockRagBroker, mockVectorRepo
}

func TestChatCommandService_CreateChat(t *testing.T) {
	service, mockChatRepo, _, _, _, _ := setupService(t)

	userID := "user-1"
	chatReq := dto.ChatCreateRequest{
		ParticipantIDs: []string{"user-2"},
	}
	expectedChatID := "new-chat-id"

	t.Run("Success", func(t *testing.T) {
		mockChatRepo.EXPECT().CreateChat(gomock.Any()).DoAndReturn(func(chat *entity.Chat) (string, error) {
			assert.Contains(t, chat.ParticipantIDs, userID)
			assert.Contains(t, chat.ParticipantIDs, "user-2")
			return expectedChatID, nil
		})

		chatID, err := service.CreateChat(chatReq, userID)

		assert.NoError(t, err)
		assert.Equal(t, expectedChatID, chatID)
	})

	t.Run("Error from repository", func(t *testing.T) {
		mockChatRepo.EXPECT().CreateChat(gomock.Any()).Return("", errors.New("db error"))

		_, err := service.CreateChat(chatReq, userID)

		assert.Error(t, err)
	})
}

func TestChatCommandService_UpdateChat(t *testing.T) {
	service, mockChatRepo, _, _, _, _ := setupService(t)

	chatID := "chat-1"
	newTitle := "New Title"
	updateReq := dto.ChatUpdateRequest{
		Title: &newTitle,
	}

	t.Run("Success", func(t *testing.T) {
		existingChat := &entity.Chat{ID: chatID, Title: new(string)}
		mockChatRepo.EXPECT().FindChatByID(chatID).Return(existingChat, nil)
		mockChatRepo.EXPECT().UpdateChat(gomock.Any()).DoAndReturn(func(chat *entity.Chat) error {
			assert.Equal(t, newTitle, *chat.Title)
			return nil
		})

		err := service.UpdateChat(chatID, updateReq)
		assert.NoError(t, err)
	})

	t.Run("Error on FindChatByID", func(t *testing.T) {
		mockChatRepo.EXPECT().FindChatByID(chatID).Return(nil, errors.New("not found"))

		err := service.UpdateChat(chatID, updateReq)
		assert.Error(t, err)
	})

	t.Run("Error on UpdateChat", func(t *testing.T) {
		existingChat := &entity.Chat{ID: chatID}
		mockChatRepo.EXPECT().FindChatByID(chatID).Return(existingChat, nil)
		mockChatRepo.EXPECT().UpdateChat(gomock.Any()).Return(errors.New("update failed"))

		err := service.UpdateChat(chatID, updateReq)
		assert.Error(t, err)
	})
}

func TestChatCommandService_SendMessage(t *testing.T) {
	service, mockChatRepo, mockParticipantRepo, mockEventBroker, mockRagBroker, _ := setupService(t)

	chatID := "chat-1"
	userID := "user-1"
	aiCoachID := "ai-coach-1"
	otherUserID := "user-2"

	t.Run("Success - send question to event broker and rag broker", func(t *testing.T) {
		req := dto.ChatItemRequest{ParticipantID: userID, Content: "Hello?", Type: "question"}

		// Mock for getOtherParticipants and getOtherParticipantIDs
		mockChatRepo.EXPECT().FindParticipantIDsByChatID(chatID).Return([]string{userID, aiCoachID, otherUserID}, nil).Times(2)
		mockParticipantRepo.EXPECT().FindByIDs([]string{userID, aiCoachID, otherUserID}).Return([]*entity.Participant{
			{ID: userID, Role: "user"},
			{ID: aiCoachID, Role: "ai_coach"},
			{ID: otherUserID, Role: "user"},
		}, nil)

		// Expect event to RAG broker
		mockRagBroker.EXPECT().PublishChatEvent(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, event dto.ChatEvent) error {
			assert.Equal(t, "rag_request", event.Type)
			assert.Equal(t, userID, event.From)
			assert.Equal(t, []string{aiCoachID}, event.To)
			return nil
		})

		// Expect event to general event broker
		mockEventBroker.EXPECT().PublishChatEvent(gomock.Any(), gomock.Any()).DoAndReturn(func(_ context.Context, event dto.ChatEvent) error {
			assert.Equal(t, "message", event.Type)
			assert.Equal(t, userID, event.From)
			assert.Contains(t, event.To, aiCoachID)
			assert.Contains(t, event.To, otherUserID)
			assert.NotContains(t, event.To, userID)
			payload, ok := event.Payload.(dto.ChatItem)
			assert.True(t, ok)
			assert.Equal(t, "question", payload.Type)
			return nil
		})

		err := service.SendMessage(chatID, req)
		assert.NoError(t, err)
	})
}

func TestChatCommandService_SaveMessage(t *testing.T) {
	service, mockChatRepo, _, _, _, mockVectorRepo := setupService(t)

	t.Run("Success - save question", func(t *testing.T) {
		item := dto.ChatItem{ID: "q-1", ChatID: "chat-1", Type: "question"}
		mockChatRepo.EXPECT().AddQuestion(item.ChatID, gomock.Any()).Return(nil)

		err := service.SaveMessage(item)
		assert.NoError(t, err)
	})

	t.Run("Success - save answer", func(t *testing.T) {
		questionID := "q-1"
		item := dto.ChatItem{ID: "a-1", ChatID: "chat-1", Type: "answer", QuestionID: &questionID}
		mockChatRepo.EXPECT().AddAnswer(item.ChatID, gomock.Any()).Return(nil)

		err := service.SaveMessage(item)
		assert.NoError(t, err)
	})

	t.Run("Success - save ai_answer and save to vector store", func(t *testing.T) {
		questionID := "q-1"
		item := dto.ChatItem{
			ID:         "ai-a-1",
			ChatID:     "chat-1",
			Type:       "ai_answer",
			QuestionID: &questionID,
			Content:    "This is an AI answer.",
		}
		mockChatRepo.EXPECT().GetQuestionContent("q-1").Return("The question?", nil)
		mockVectorRepo.EXPECT().SaveQAPair("chat-1", "The question?", "This is an AI answer.", "ai-a-1").Return(nil)
		mockChatRepo.EXPECT().AddAnswer("chat-1", gomock.Any()).Return(nil)

		err := service.SaveMessage(item)
		assert.NoError(t, err)
	})

	t.Run("Success - save ai_answer with GetQuestionContent error", func(t *testing.T) {
		item := dto.ChatItem{Type: "ai_answer", ChatID: "chat-1"}
		mockChatRepo.EXPECT().GetQuestionContent(gomock.Any()).Return("", errors.New("db error"))
		mockVectorRepo.EXPECT().SaveQAPair(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(0)
		mockChatRepo.EXPECT().AddAnswer("chat-1", gomock.Any()).Return(nil)

		err := service.SaveMessage(item)
		assert.NoError(t, err)
	})

	t.Run("Error on AddQuestion", func(t *testing.T) {
		item := dto.ChatItem{Type: "question", ChatID: "chat-1"}
		mockChatRepo.EXPECT().AddQuestion(gomock.Any(), gomock.Any()).Return(errors.New("db error"))

		err := service.SaveMessage(item)
		assert.Error(t, err)
	})

	t.Run("Error on AddAnswer", func(t *testing.T) {
		item := dto.ChatItem{Type: "answer", ChatID: "chat-1"}
		mockChatRepo.EXPECT().AddAnswer(gomock.Any(), gomock.Any()).Return(errors.New("db error"))

		err := service.SaveMessage(item)
		assert.Error(t, err)
	})
}