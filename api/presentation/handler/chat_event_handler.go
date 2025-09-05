package handler

import (
	"api/application/dto"
	"api/application/service/command"
	"encoding/json"
)

func ChatEventHandler(event dto.ChatEvent, service *command.ChatCommandService) error {
	switch event.Type {
	case "question":
		payloadBytes, err := json.Marshal(event.Payload)
		if err != nil {
			return err
		}
		var q dto.QuestionResponse
		if err := json.Unmarshal(payloadBytes, &q); err != nil {
			return err
		}
		return service.SaveQuestion(q)
	case "answer":
		payloadBytes, err := json.Marshal(event.Payload)
		if err != nil {
			return err
		}
		var a dto.AnswerResponse
		if err := json.Unmarshal(payloadBytes, &a); err != nil {
			return err
		}
		return service.SaveAnswer(a)
	default:
		return nil
	}
}
