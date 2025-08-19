package dto

type ParticipantUpdateRequest struct {
	ID       string  `json:"id"`
	Name     *string `json:"name,omitempty"`
	Role     *string `json:"role,omitempty"`
	Sports   []string `json:"sports,omitempty"`
	IconURL  *string `json:"icon_url,omitempty"`
}
