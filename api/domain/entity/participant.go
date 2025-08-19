package entity

type Participant struct {
	ID     string
	Name   string
	Email  string
	Role   string
	Sports []string
	IconURL *string
}
