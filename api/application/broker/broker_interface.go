package broker

import (
	"context"
	"api/application/dto"
)

type ChatEventPublisher interface {
	PublishChatEvent(ctx context.Context, event dto.ChatEvent) error
}