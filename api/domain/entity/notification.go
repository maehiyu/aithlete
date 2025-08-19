package entity

type Notification struct {
	ID        string
	UserID    string
	ChatID    string
	Type      string
	Content   string
	CreatedAt string // datetime型推奨
	Read      bool
}
