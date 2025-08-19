package dto

type ChatCreateRequest struct {
	Title          *string  `json:"title,omitempty"`
	ParticipantIDs []string `json:"participant_ids"`
}
