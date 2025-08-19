package dto

import "time"

type ChatSummaryResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	LastActiveAt time.Time `json:"last_active_at"`
	LatestQA     string `json:"latest_qa"`
	Opponent     OpponentResponse `json:"opponent"`
}

type OpponentResponse struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	IconURL string `json:"icon_url"`
}
