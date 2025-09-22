package query

import (
	"api/application/dto"
	"api/application/query"
)

type ParticipantQueryService struct {
	ParticipantQuery query.ParticipantQueryInterface
}

func NewParticipantQueryService(qs query.ParticipantQueryInterface) *ParticipantQueryService {
	return &ParticipantQueryService{ParticipantQuery: qs}
}
	
func (s *ParticipantQueryService) GetParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error) {
	participants, err := s.ParticipantQuery.FindParticipantsByChatID(chatID)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ParticipantResponse, len(participants))
	for i, p := range participants {
		response[i] = *dto.ParticipantEntityToResponse(&p)
	}
	return response, nil
}

func (s *ParticipantQueryService) GetParticipantByID(participantID string) (*dto.ParticipantResponse, error) {
	participant, err := s.ParticipantQuery.FindParticipantByID(participantID)
	if err != nil {
		return nil, err
	}
	response := dto.ParticipantEntityToResponse(participant)
	return response, nil
}

func (s *ParticipantQueryService) GetCoachesBySport(sport string) ([]dto.ParticipantResponse, error) {
	participants, err := s.ParticipantQuery.FindCoachesBySport(sport)
	if err != nil {
		return nil, err
	}
	response := make([]dto.ParticipantResponse, len(participants))
	for i, p := range participants {
		response[i] = *dto.ParticipantEntityToResponse(&p)
	}
	return response, nil
}