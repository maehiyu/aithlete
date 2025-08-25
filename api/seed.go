package main

import (
	"api/domain/entity"
	"encoding/json"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	// Participants
	participants := []entity.Participant{
		{ID: "user1", Name: "ユーザー花子", Email: "user1@example.com", Role: "user", Sports: []string{"サッカー", "バスケットボール"}},
		{ID: "coach1", Name: "コーチ太郎", Email: "coach1@example.com", Role: "coach", Sports: []string{"野球"}},
		{ID: "coach2", Name: "コーチ二郎", Email: "coach2@example.com", Role: "coach", Sports: []string{"テニス"}},
	}
	for _, p := range participants {
		db.FirstOrCreate(&p, entity.Participant{ID: p.ID})
	}

	// Chats
	title1 := "ユーザーとコーチのチャット"
	title2 := "コーチからのアドバイス"
	chats := []entity.Chat{
		{
			ID:             "chat1",
			Title:          &title1,
			ParticipantIDs: []string{"user1", "coach1"},
			Questions: []entity.Question{
				{ID: "q1", ChatID: "chat1", Content: "コーチへの質問です", ParticipantID: "user1", CreatedAt: time.Now()},
			},
			Answers: []entity.Answer{
				{ID: "a1", QuestionID: "q1", ParticipantID: "coach1", Content: "コーチからの回答です", CreatedAt: time.Now()},
			},
		},
		{
			ID:             "chat2",
			Title:          &title2,
			ParticipantIDs: []string{"user1", "coach2"},
			Questions: []entity.Question{
				{ID: "q2", ChatID: "chat2", Content: "アドバイス内容です", ParticipantID: "coach2", CreatedAt: time.Now()},
			},
			Answers: []entity.Answer{
				{ID: "a2", QuestionID: "q2", ParticipantID: "user1", Content: "ユーザーからの返信です", CreatedAt: time.Now()},
			},
		},
	}
	for _, c := range chats {
		db.FirstOrCreate(&c, entity.Chat{ID: c.ID})
	}
}

func toJSON(ids []string) datatypes.JSON {
	b, _ := json.Marshal(ids)
	return datatypes.JSON(b)
}
