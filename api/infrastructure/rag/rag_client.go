package rag

import (
	"api/application/dto"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type RAGClient struct {
	Endpoint string
}

func NewRAGClient() *RAGClient {
	endpoint := os.Getenv("RAG_SERVER_URL")
	if endpoint == "" {
		endpoint = "http://rag:9200/rag/query"
	}
	return &RAGClient{Endpoint: endpoint}
}

func (c *RAGClient) CallRAGServer(chatID, questionContent, aiID, questionID, token string) (string, error) {
	reqBody := dto.RagQueryRequest{
		Question:      questionContent,
		ChatID:        chatID,
		QuestionID:    questionID,
		ParticipantID: aiID,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	bearerToken := token
	if !strings.HasPrefix(token, "Bearer ") {
		bearerToken = "Bearer " + token
	}
	req.Header.Set("Authorization", bearerToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("RAG server returned status: %s", resp.Status)
	}
	var result struct {
		Answer string `json:"answer"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Answer, nil
}
