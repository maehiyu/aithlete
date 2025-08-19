package query

import (
	"api/application/dto"
	"api/application/query"
)

type ParticipantQueryService struct {
	queryService query.ParticipantQueryServiceInterface
}

func NewParticipantQueryService(qs query.ParticipantQueryServiceInterface) *ParticipantQueryService {
	return &ParticipantQueryService{queryService: qs}
}

func (s *ParticipantQueryService) FindParticipantsByChatID(chatID string) ([]dto.ParticipantResponse, error) {
	participants, err := s.queryService.FindParticipantsByChatID(chatID)
	if err != nil {
		return nil, err
	}
	return participants, nil
}
