package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"encoding/json"

	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var cognitoKeyfunc keyfunc.Keyfunc

func init() {
	region := os.Getenv("COGNITO_REGION")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	if region == "" || userPoolID == "" {
		log.Println("COGNITO_REGION and COGNITO_USER_POOL_ID must be set in environment variables")
		return
	}
	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	ctx := context.Background()
	var err error
	cognitoKeyfunc, err = keyfunc.NewDefaultCtx(ctx, []string{jwksURL})
	if err != nil {
		log.Fatalf("failed to create keyfunc.Keyfunc: %v", err)
	}
}

func extractUserIDFromJWT(r *http.Request) (string, error) {
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {

		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			return "", fmt.Errorf("missing or invalid Authorization header and no token param")
		}
		tokenString = strings.TrimPrefix(auth, "Bearer ")
	}
	token, err := jwt.Parse(tokenString, cognitoKeyfunc.Keyfunc)
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid JWT: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid JWT claims")
	}
	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", fmt.Errorf("sub not found in token")
	}
	return sub, nil
}

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
	clients map[string]map[*Client]bool 
	mu      sync.Mutex
}

func NewHub() *Hub {
	return &Hub{clients: make(map[string]map[*Client]bool)}
}

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

	go func() {
		for msg := range c.send {
			err := c.conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				break
			}
		}
		c.conn.Close()
	}()
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
		userId, err := extractUserIDFromJWT(r)
		if err != nil {
			log.Println("JWT error:", err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

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
			for {
				_, _, err := conn.ReadMessage()
				if err != nil {
					break
				}
			}
		}()
	}
}

func main() {
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

	go func() {
		ctx := context.Background()
		sub := rdb.Subscribe(ctx, "chat_events","chat_stream")
		ch := sub.Channel()
		for msg := range ch {
			var event ChatEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err == nil {
				if !(event.Type == "chat_event" && event.From == "ai_coach") {
					hub.SendToUsers([]byte(msg.Payload), event.To)
				}
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
