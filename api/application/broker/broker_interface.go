//go:generate mockgen -source=broker_interface.go -destination=mocks/mock_chat_event_broker.go -package=mocks ChatEventBroker
package broker

import (
	"api/application/dto"
	"context"
)

type ChatEventBroker interface {
	PublishChatEvent(ctx context.Context, event dto.ChatEvent) error
	SubscribeChatEvent(ctx context.Context, handler func(dto.ChatEvent) error) error
}
