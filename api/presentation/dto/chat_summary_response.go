package dto

type ChatSummaryResponse struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	LastActiveAt string `json:"last_active_at"`
	LatestQA     string `json:"latest_qa"`
}
