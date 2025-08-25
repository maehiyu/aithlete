package command

import "api/domain/entity"

type MockParticipantRepository struct {
	FindByIDFunc func(id string) (*entity.Participant, error)
	CreateFunc   func(participant *entity.Participant) (*entity.Participant, error)
	UpdateFunc   func(participant *entity.Participant) (*entity.Participant, error)
}

func NewMockParticipantRepository() *MockParticipantRepository {
	return &MockParticipantRepository{
		FindByIDFunc: func(id string) (*entity.Participant, error) { panic("not used") },
		CreateFunc:   func(participant *entity.Participant) (*entity.Participant, error) { panic("not used") },
		UpdateFunc:   func(participant *entity.Participant) (*entity.Participant, error) { panic("not used") },
	}
}

func (m *MockParticipantRepository) FindByID(id string) (*entity.Participant, error) {
	return m.FindByIDFunc(id)
}

func (m *MockParticipantRepository) Create(participant *entity.Participant) (*entity.Participant, error) {
	return m.CreateFunc(participant)
}

func (m *MockParticipantRepository) Update(participant *entity.Participant) (*entity.Participant, error) {
	return m.UpdateFunc(participant)
}
