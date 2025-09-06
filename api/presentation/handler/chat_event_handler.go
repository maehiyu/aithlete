package handler

import (
	"api/application/dto"
	"api/application/service/command"
	"encoding/json"
	"fmt"
)

func ChatEventHandler(event dto.ChatEvent, service *command.ChatCommandService) error {
	fmt.Printf("ChatEventHandler event: %v\n", event)
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
	case "answer", "ai_answer":
		payloadBytes, err := json.Marshal(event.Payload)
		if err != nil {
			return err
		}
		var a dto.AnswerResponse
		fmt.Printf("%s", payloadBytes)
		if err := json.Unmarshal(payloadBytes, &a); err != nil {
			return err
		}
		fmt.Printf("ChatEventHandler answer: chat_id=%v, question_id=%v, participant_id=%v\n", a.ChatID, a.QuestionID, a.ParticipantID)
		return service.SaveAnswer(a)
	default:
		return nil
	}
}
