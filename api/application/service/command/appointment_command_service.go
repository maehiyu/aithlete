package command

import (
	"api/application/dto"
	"api/domain/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type AppointmentCommandService struct {
	repo repository.AppointmentRepositoryInterface
}

func NewAppointmentCommandService(repo repository.AppointmentRepositoryInterface) *AppointmentCommandService {
	return &AppointmentCommandService{repo: repo}
}

func (s *AppointmentCommandService) CreateAppointment(ctx context.Context, req dto.AppointmentCreateRequest) (string, error) {
	now := time.Now()
	id := uuid.NewString()

	appointment := dto.AppointmentCreateRequestToEntity(req, id, now)

	id, err := s.repo.Create(ctx, appointment)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *AppointmentCommandService) UpdateAppointment(ctx context.Context, id string, req dto.AppointmentUpdateRequest) error {
	appointment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	dto.AppointmentUpdateRequestToEntity(appointment, req)

	if err := s.repo.Update(ctx, appointment); err != nil {
		return err
	}

	return nil
}

func (s *AppointmentCommandService) DeleteAppointment(ctx context.Context, id string) error {
	// TODO: 削除権限のチェックなどをここに実装

	return s.repo.Delete(ctx, id)
}
