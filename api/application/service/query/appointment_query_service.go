package query

import (
	"api/application/dto"
	"api/application/query"
	"context"
)

// AppointmentQueryService は予約に関する読み取り系のユースケースを実装します。
type AppointmentQueryService struct {
	query query.AppointmentQueryInterface
}

// NewAppointmentQueryService はAppointmentQueryServiceの新しいインスタンスを生成します。
func NewAppointmentQueryService(query query.AppointmentQueryInterface) *AppointmentQueryService {
	return &AppointmentQueryService{query: query}
}

// GetByID はIDで予約を取得します。
func (s *AppointmentQueryService) GetByID(ctx context.Context, id string) (*dto.AppointmentResponse, error) {
	return s.query.GetByID(ctx, id)
}

// GetByChatID はチャットIDで予約のリストを取得します。
func (s *AppointmentQueryService) GetByChatID(ctx context.Context, chatID string) ([]*dto.AppointmentResponse, error) {
	return s.query.GetByChatID(ctx, chatID)
}

// GetByUserID はユーザーIDで予約のリストを取得します。
func (s *AppointmentQueryService) GetByUserID(ctx context.Context, userID string) ([]*dto.AppointmentResponse, error) {
	return s.query.GetByUserID(ctx, userID)
}

// GetByCoachID はコーチIDで予約のリストを取得します。
func (s *AppointmentQueryService) GetByCoachID(ctx context.Context, coachID string) ([]*dto.AppointmentResponse, error) {
	return s.query.GetByCoachID(ctx, coachID)
}
