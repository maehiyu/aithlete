package ai

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Request struct {
	Input string `json:"input"`
}

type Response struct {
	Output string `json:"output"`
}

func coachHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// ここでAIコーチング処理を行う（今回はエコー返し）
	resp := Response{Output: "AIコーチング応答: " + req.Input}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/coach", coachHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9100"
	}
	log.Printf("AI Coaching Server listening on :%s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
