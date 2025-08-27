package dto

type ChatEvent struct {
	ID        string    `json:"id"`
	ChatID    string    `json:"chat_id"`
	Type      string    `json:"type"`
	From      string    `json:"from"`
	To        []string  `json:"to"`
	Timestamp int64     `json:"timestamp"`
	Payload   any       `json:"payload"`
}
