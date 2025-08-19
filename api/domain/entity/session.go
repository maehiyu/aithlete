package entity

type Session struct {
	ID            string
	ParticipantID string
	ChatID        string
	StartedAt     string
	LastActiveAt  string
	Status        string 
	IPAddress     string
}
