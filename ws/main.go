package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

// --- domain ---
type ChatEvent struct {
	ID        string      `json:"id"`
	ChatID    string      `json:"chat_id"`
	Type      string      `json:"type"`
	From      string      `json:"from"`
	To        []string    `json:"to"`
	Timestamp int64       `json:"timestamp"`
	Payload   interface{} `json:"payload"`
}

// --- infrastructure ---
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	userId string
}


type Hub struct {
	clients map[string]map[*Client]bool // userId -> set of clients
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{clients: make(map[string]map[*Client]bool)}
}

import "encoding/json"

func (h *Hub) SendToUsers(msg []byte, to []string) {
       h.mu.Lock()
       defer h.mu.Unlock()
       for _, userId := range to {
	       for c := range h.clients[userId] {
		       select {
		       case c.send <- msg:
		       default:
			       close(c.send)
			       delete(h.clients[userId], c)
		       }
	       }
       }
}

func (h *Hub) AddClient(c *Client) {
       h.mu.Lock()
       if h.clients[c.userId] == nil {
	       h.clients[c.userId] = make(map[*Client]bool)
       }
       h.clients[c.userId][c] = true
       h.mu.Unlock()
}

func (h *Hub) RemoveClient(c *Client) {
       h.mu.Lock()
       if m, ok := h.clients[c.userId]; ok {
	       if _, ok2 := m[c]; ok2 {
		       delete(m, c)
		       close(c.send)
		       if len(m) == 0 {
			       delete(h.clients, c.userId)
		       }
	       }
       }
       h.mu.Unlock()
}

func wsHandler(hub *Hub) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       // JWTからuserIdを取得（今回はダミーでuser1固定）
	       // authHeader := r.Header.Get("Authorization")
	       userId := "user1" // ←ここをJWTデコードに置き換えればOK

	       conn, err := upgrader.Upgrade(w, r, nil)
	       if err != nil {
		       log.Println("upgrade error:", err)
		       return
	       }
	       client := &Client{conn: conn, send: make(chan []byte, 256), userId: userId}
	       hub.AddClient(client)

	       go func() {
		       defer func() {
			       hub.RemoveClient(client)
			       conn.Close()
		       }()
		       for msg := range client.send {
			       if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				       break
			       }
		       }
	       }()
       }
}

func main() {
	// Redis接続
	redisHost := os.Getenv("BROKER_HOST")
	if redisHost == "" {
		redisHost = "broker"
	}
	redisPort := os.Getenv("BROKER_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisPassword := os.Getenv("BROKER_PASSWORD")
	redisDB := 0
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       redisDB,
	})

	hub := NewHub()

       // Redis購読goroutine
       go func() {
	       ctx := context.Background()
	       sub := rdb.Subscribe(ctx, "chat_events")
	       ch := sub.Channel()
	       for msg := range ch {
		       // ChatEventのToだけに送信
		       var event ChatEvent
		       if err := json.Unmarshal([]byte(msg.Payload), &event); err == nil {
			       hub.SendToUsers([]byte(msg.Payload), event.To)
		       }
	       }
       }()

	http.HandleFunc("/ws", wsHandler(hub))
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	fmt.Println("WebSocket server started on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
