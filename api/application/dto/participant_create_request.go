package dto

type ParticipantCreateRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Sports  []string `json:"sports,omitempty"`
	IconURL *string `json:"icon_url"`
}
