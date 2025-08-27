package broker

import (
	"api/application/dto"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type RedisChatEventPublisher struct {
	client *redis.Client
	topic  string
}

func NewRedisChatEventPublisher(client *redis.Client, topic string) *RedisChatEventPublisher {
	return &RedisChatEventPublisher{client: client, topic: topic}
}

func (r *RedisChatEventPublisher) PublishChatEvent(ctx context.Context, event dto.ChatEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return r.client.Publish(ctx, r.topic, payload).Err()
}
