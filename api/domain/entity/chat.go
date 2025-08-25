package entity

import (
	"time"
)

type Chat struct {
	ID             string
	StartedAt      time.Time
	LastActiveAt   time.Time
	Title          *string
	Questions      []Question
	Answers        []Answer
	ParticipantIDs []string
}

type Question struct {
	ID            string
	ChatID        string
	ParticipantID string
	Content       string
	Attachments   []Attachment
	CreatedAt     time.Time
}

type Answer struct {
	ID            string
	ChatID        string
	QuestionID    string
	ParticipantID string
	Content       string
	Attachments   []Attachment
	CreatedAt     time.Time
}

type Attachment struct {
	ID         string
	Type       AttachmentType
	URL        string
	Thumbnail  *string
	PoseID     *string
	Pose       *PoseData
	Meta       string
	OriginalID *string
	Original   *Attachment
	QuestionID *string
	AnswerID   *string
}

type AttachmentType string

const (
	AttachmentTypeVideo         AttachmentType = "video"
	AttachmentTypeSkeletonVideo AttachmentType = "skeleton_video"
	AttachmentTypeImage         AttachmentType = "image"
	AttachmentTypePDF           AttachmentType = "pdf"
)

type PoseData struct {
	ID             string
	ParticipantIDs []string
	Score          float64
}
