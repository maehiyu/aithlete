package command

import (
	"api/application/dto"
	"api/domain/entity"
	"api/domain/repository/mocks" // gomockで生成されるモックのパスを想定
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestParticipantCommandService_CreateParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockParticipantRepositoryInterface(ctrl)
	service := NewParticipantCommandService(mockRepo)

	testCases := []struct {
		name          string
		req           dto.ParticipantCreateRequest
		userID        string
		setupMock     func(mock *mocks.MockParticipantRepositoryInterface)
		expectedID    string
		expectErr     bool
		expectedErr   error
	}{
		{
			name: "Success - Create normal user",
			req: dto.ParticipantCreateRequest{
				Name: "Test User",
				Role: "user",
			},
			userID: "user-id-123",
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().Create(gomock.Any()).DoAndReturn(func(p *entity.Participant) (string, error) {
					assert.Equal(t, "Test User", p.Name)
					assert.Equal(t, "user", p.Role)
					assert.Equal(t, "user-id-123", p.ID)
					return "user-id-123", nil
				})
			},
			expectedID: "user-id-123",
			expectErr:  false,
		},
		{
			name: "Success - Create AI coach",
			req: dto.ParticipantCreateRequest{
				Name: "AI Coach",
				Role: "ai_coach",
			},
			userID: "some-user-id", // This should be ignored
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().Create(gomock.Any()).DoAndReturn(func(p *entity.Participant) (string, error) {
					assert.Equal(t, "AI Coach", p.Name)
					assert.Equal(t, "ai_coach", p.Role)
					assert.NotEqual(t, "some-user-id", p.ID) // Should be a new UUID
					return "generated-ai-id", nil
				})
			},
			expectedID: "generated-ai-id",
			expectErr:  false,
		},
		{
			name: "Error - Repository create fails",
			req: dto.ParticipantCreateRequest{
				Name: "Fail User",
				Role: "user",
			},
			userID: "user-id-fail",
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().Create(gomock.Any()).Return("", errors.New("db error"))
			},
			expectedID: "",
			expectErr:  true,
			expectedErr: errors.New("db error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock(mockRepo)
			id, err := service.CreateParticipant(tc.req, tc.userID)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedID, id)
			}
		})
	}
}

func TestParticipantCommandService_UpdateParticipant(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockParticipantRepositoryInterface(ctrl)
	service := NewParticipantCommandService(mockRepo)

	participantID := "participant-to-update"
	existingParticipant := &entity.Participant{
		ID:   participantID,
		Name: "Old Name",
		Role: "user",
	}
	newName := "New Name"

	testCases := []struct {
		name        string
		participantID string
		req         dto.ParticipantUpdateRequest
		setupMock   func(mock *mocks.MockParticipantRepositoryInterface)
		expectErr   bool
		expectedErr error
	}{
		{
			name:        "Success - Update participant",
			participantID: participantID,
			req: dto.ParticipantUpdateRequest{
				Name: &newName,
			},
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().FindByID(participantID).Return(existingParticipant, nil)
				mock.EXPECT().Update(gomock.Any()).DoAndReturn(func(p *entity.Participant) error {
					assert.Equal(t, newName, p.Name)
					return nil
				})
			},
			expectErr: false,
		},
		{
			name:        "Error - FindByID fails",
			participantID: participantID,
			req: dto.ParticipantUpdateRequest{
				Name: &newName,
			},
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().FindByID(participantID).Return(nil, errors.New("not found"))
			},
			expectErr:   true,
			expectedErr: errors.New("not found"),
		},
		{
			name:        "Error - Update fails",
			participantID: participantID,
			req: dto.ParticipantUpdateRequest{
				Name: &newName,
			},
			setupMock: func(mock *mocks.MockParticipantRepositoryInterface) {
				mock.EXPECT().FindByID(participantID).Return(existingParticipant, nil)
				mock.EXPECT().Update(gomock.Any()).Return(errors.New("db update error"))
			},
			expectErr:   true,
			expectedErr: errors.New("db update error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMock(mockRepo)
			err := service.UpdateParticipant(tc.participantID, tc.req)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}