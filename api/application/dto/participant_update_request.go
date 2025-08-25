package dto

type ParticipantUpdateRequest struct {
	Name     *string `json:"name,omitempty"`
	Email	*string `json:"email,omitempty"`
	Role     *string `json:"role,omitempty"`
	Sports   []string `json:"sports,omitempty"`
	IconURL  *string `json:"icon_url,omitempty"`
}
