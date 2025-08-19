package entity

import "time"

type Chat struct {
	ID            string `gorm:"primaryKey"`
	StartedAt     time.Time
	LastActiveAt  time.Time
	Title         *string
	Questions     []Question `gorm:"foreignKey:ChatID"`
	Answers       []Answer   `gorm:"foreignKey:ChatID"`
	ParticipantIDs []string   `gorm:"type:text[]"`
}

type Question struct {
	ID            string `gorm:"primaryKey"`
	ChatID        string // 外部キー
	ParticipantID string
	Content       string
	Attachments   []Attachment `gorm:"foreignKey:QuestionID"`
	CreatedAt     time.Time
}

type Answer struct {
	ID            string `gorm:"primaryKey"`
	ChatID        string // 外部キー
	QuestionID    string
	ParticipantID string
	Content       string
	Attachments   []Attachment `gorm:"foreignKey:AnswerID"`
	CreatedAt     time.Time
}

type Attachment struct {
	ID         string `gorm:"primaryKey"`
	Type       AttachmentType
	URL        string
	Thumbnail  *string
	PoseID     *string
	Pose       *PoseData   `gorm:"foreignKey:PoseID"`
	Meta       string      `gorm:"type:jsonb"` // JSON文字列で格納
	OriginalID *string     // スケルトン動画の場合、元動画のID
	Original   *Attachment `gorm:"foreignKey:OriginalID"`
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
	ID        string `gorm:"primaryKey"`
	Keypoints string
	Score     float64
}
