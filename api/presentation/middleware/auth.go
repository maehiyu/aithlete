package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// ContextKeyUserID is the key used to store userID in Gin context
const ContextKeyUserID = "userId"

// AuthMiddleware extracts userId from JWT (dummy implementation)
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		// 本来はtokenを検証しuserIdを抽出する処理を実装
		userId := token // ダミー: token自体をuserIdとみなす
		if userId == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT"})
			return
		}
		c.Set(ContextKeyUserID, userId)
		c.Next()
	}
}
