package query

import (
	"api/application/dto"
	"api/application/query"
	"context"
)

type ParticipantQueryService struct {
	ParticipantQuery query.ParticipantQueryInterface
}

func NewParticipantQueryService(qs query.ParticipantQueryInterface) *ParticipantQueryService {
	return &ParticipantQueryService{ParticipantQuery: qs}
}
	
func (s *ParticipantQueryService) GetParticipantsByChatID(ctx context.Context, chatID string) ([]dto.ParticipantResponse, error) {
	participants, err := s.ParticipantQuery.FindParticipantsByChatID(ctx, chatID)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ParticipantResponse, len(participants))
	for i, p := range participants {
		response[i] = *dto.ParticipantEntityToResponse(&p)
	}
	return response, nil
}

func (s *ParticipantQueryService) GetParticipantByID(ctx context.Context, participantID string) (*dto.ParticipantResponse, error) {
	participant, err := s.ParticipantQuery.FindParticipantByID(ctx, participantID)
	if err != nil {
		return nil, err
	}
	response := dto.ParticipantEntityToResponse(participant)
	return response, nil
}

func (s *ParticipantQueryService) GetCoachesBySport(ctx context.Context, sport string) ([]dto.ParticipantResponse, error) {
	participants, err := s.ParticipantQuery.FindCoachesBySport(ctx, sport)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ParticipantResponse, len(participants))
	for i, p := range participants {
		response[i] = *dto.ParticipantEntityToResponse(&p)
	}
	return response, nil
}