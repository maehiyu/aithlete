package repository

import (
	"api/domain/entity"
	"api/domain/repository"
)

type ParticipantRepositoryImpl struct {
	
}

func (r *ParticipantRepositoryImpl) FindByID(participantID string) (*entity.Participant, error) {
	// 実装例: DBからparticipantIDでParticipantを取得
	return &entity.Participant{}, nil
}

func (r *ParticipantRepositoryImpl) Create(participant *entity.Participant) (*entity.Participant, error) {
	// 実装例: DBにParticipantを保存
	return participant, nil
}

func (r *ParticipantRepositoryImpl) Update(participant *entity.Participant) (*entity.Participant, error) {
	// 実装例: DBのParticipantを更新
	return participant, nil
}

var _ repository.ParticipantRepositoryInterface = (*ParticipantRepositoryImpl)(nil)