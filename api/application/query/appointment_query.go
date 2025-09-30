//go:generate mockgen -source=appointment_query.go -destination=mocks/mock_appointment_query.go -package=mocks AppointmentQueryInterface
package query

import (
	"api/application/dto"
	"context"
)

type AppointmentQueryInterface interface {
	// GetByID はIDで単一の予約を取得します（参加者情報も含む）
	GetByID(ctx context.Context, id string) (*dto.AppointmentResponse, error)

	// ListByUserID は特定のユーザーが参加している予約のリストを取得します
	ListByUserID(ctx context.Context, userID string) ([]*dto.AppointmentResponse, error)

	// ListByCoachID は特定のコーチが開催する予約のリストを取得します
	ListByCoachID(ctx context.Context, coachID string) ([]*dto.AppointmentResponse, error)

	// ListByChatID は特定のチャットに関連する予約のリストを取得します
	ListByChatID(ctx context.Context, chatID string) ([]*dto.AppointmentResponse, error)
}