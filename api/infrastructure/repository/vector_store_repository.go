package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type VectorStoreRepository struct {
	WeaviateEndpoint  string // 例: "http://weaviate:8080"
	WeaviateClass     string // 例: "QAPair"
	EmbeddingEndpoint string // 例: "http://embedding:8001/embed"
}

func NewVectorStoreRepository(weaviateEndpoint, weaviateClass, embeddingEndpoint string) *VectorStoreRepository {
	return &VectorStoreRepository{
		WeaviateEndpoint:  weaviateEndpoint,
		WeaviateClass:     weaviateClass,
		EmbeddingEndpoint: embeddingEndpoint,
	}
}

func (v *VectorStoreRepository) SaveQAPair(chatID, question, answer, answerID string) error {
	txt := fmt.Sprintf("%s %s", question, answer)
	payload := map[string]any{
		"texts": []string{txt},
	}
	body, _ := json.Marshal(payload)

	// デバッグログ追加
	fmt.Printf("[VECTOR_STORE] Sending embedding request: %s\n", string(body))

	resp, err := http.Post(v.EmbeddingEndpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		// エラーレスポンスの詳細を取得
		respBody := make([]byte, 1024)
		n, _ := resp.Body.Read(respBody)
		fmt.Printf("[VECTOR_STORE] Embedding API error: %s, Response: %s\n", resp.Status, string(respBody[:n]))
		return fmt.Errorf("embedding api error: %s", resp.Status)
	}
	var embResp struct {
		Embeddings [][]float32 `json:"embeddings"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&embResp); err != nil {
		return err
	}
	if len(embResp.Embeddings) == 0 {
		return fmt.Errorf("no embedding returned")
	}

	// 2. Weaviateに保存
	weaviatePayload := map[string]interface{}{
		"class": v.WeaviateClass,
		"id":    answerID,
		"properties": map[string]interface{}{
			"chat_id":  chatID,
			"question": question,
			"answer":   answer,
		},
		"vector": embResp.Embeddings[0],
	}
	wbody, _ := json.Marshal(weaviatePayload)
	url := fmt.Sprintf("%s/v1/objects", v.WeaviateEndpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(wbody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	wresp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer wresp.Body.Close()
	if wresp.StatusCode != http.StatusOK && wresp.StatusCode != http.StatusCreated {
		return fmt.Errorf("weaviate upsert failed: %s", wresp.Status)
	}
	return nil
}
