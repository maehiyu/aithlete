package broker

import (
	"api/application/dto"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisChatEventBroker struct {
	client *redis.Client
	topic  string
}

func NewRedisChatEventBroker(client *redis.Client, topic string) *RedisChatEventBroker {
	return &RedisChatEventBroker{client: client, topic: topic}
}

func (r *RedisChatEventBroker) PublishChatEvent(ctx context.Context, event dto.ChatEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, r.topic, payload).Err()
}

func (r *RedisChatEventBroker) SubscribeChatEvent(ctx context.Context, handler func(dto.ChatEvent) error) error {
	pubsub := r.client.Subscribe(ctx, r.topic)
	ch := pubsub.Channel()
	go func() {
		for msg := range ch {
			var event dto.ChatEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err == nil {
				handler(event)
			}
		}
	}()
	return nil
}
