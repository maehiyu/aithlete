package command

import (
	"api/application/dto"
	"api/domain/entity"
	"api/domain/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

// AppointmentCommandService は予約に関する書き込み系のユースケースを実装します。
type AppointmentCommandService struct {
	appointmentRepo repository.AppointmentRepositoryInterface
}

// NewAppointmentCommandService はAppointmentCommandServiceの新しいインスタンスを生成します。
func NewAppointmentCommandService(ar repository.AppointmentRepositoryInterface) *AppointmentCommandService {
	return &AppointmentCommandService{appointmentRepo: ar}
}

// CreateAppointment は新しい予約を作成し、参加者を登録します。
func (s *AppointmentCommandService) CreateAppointment(ctx context.Context, req dto.AppointmentCreateRequest, requesterID string) (string, error) {
	now := time.Now()
	id := uuid.NewString()

	// 1. 予約枠エンティティを作成
	appointment := dto.AppointmentCreateRequestToEntity(req, id, now)

	// TODO: トランザクションを開始する

	// 2. 予約枠をDBに保存
	createdAppointment, err := s.appointmentRepo.CreateAppointment(ctx, appointment)
	if err != nil {
		return "", err
	}

	// 3. 予約参加情報エンティティを作成
	appointmentParticipants := make([]*entity.AppointmentParticipant, len(req.ParticipantIDs))
	for i, pID := range req.ParticipantIDs {
		status := "needs-action" // デフォルトは「要返信」
		if pID == requesterID {
			status = "accepted" // リクエスト元（作成者）は自動で「承諾」
		}
		appointmentParticipants[i] = &entity.AppointmentParticipant{
			AppointmentID: createdAppointment.ID,
			ParticipantID: pID,
			Status:        status,
		}
	}

	// 4. 予約参加情報をDBに保存
	if err := s.appointmentRepo.CreateAppointmentParticipants(ctx, appointmentParticipants); err != nil {
		// TODO: トランザクションをロールバックする
		return "", err
	}

	// TODO: トランザクションをコミットする

	return createdAppointment.ID, nil
}

func (s *AppointmentCommandService) UpdateAppointment(ctx context.Context, id string, req dto.AppointmentUpdateRequest)  error {
	appointment, err := s.appointmentRepo.FindAppointmentByID(ctx, id)
	if err != nil {
		return err
	}

	dto.AppointmentUpdateRequestToEntity(appointment, req)

	if err := s.appointmentRepo.UpdateAppointment(ctx, appointment); err != nil {
		return err
	}

	return nil
}

// DeleteAppointment は予約を削除します。
func (s *AppointmentCommandService) DeleteAppointment(ctx context.Context, id string) error {
	// TODO: 削除権限のチェックなどをここに実装

	return s.appointmentRepo.DeleteAppointment(ctx, id)
}
