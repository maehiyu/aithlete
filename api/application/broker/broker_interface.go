package broker

import (
	"api/application/dto"
	"context"
)

type ChatEventBroker interface {
	PublishChatEvent(ctx context.Context, event dto.ChatEvent) error
	SubscribeChatEvent(ctx context.Context, handler func(dto.ChatEvent) error) error
}
