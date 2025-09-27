package handler

import (
	"api/application/dto"
	"api/application/service/command"
	"context"
	"encoding/json"
)

func ChatEventHandler(ctx context.Context, event dto.ChatEvent, service *command.ChatCommandService) error {
	payloadBytes, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}
	var item dto.ChatItem
	if err := json.Unmarshal(payloadBytes, &item); err != nil {
		return err
	}
	return service.SaveMessage(ctx, item)
}
