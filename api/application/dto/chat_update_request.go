package dto

type ChatUpdateRequest struct {
	Title         *string  `json:"title"`
	ParticipantIDs []string `json:"participant_ids"`
}
