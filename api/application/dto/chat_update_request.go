package dto

type ChatUpdateRequest struct {
	ID            string   `json:"id"`
	Title         *string  `json:"title"`
	ParticipantIDs []string `json:"participant_ids"`
}
