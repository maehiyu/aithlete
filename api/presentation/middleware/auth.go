package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const ContextKeyUserID = "userId"

var kf keyfunc.Keyfunc

func init() {
	region := os.Getenv("COGNITO_REGION")
	userPoolID := os.Getenv("COGNITO_USER_POOL_ID")
	if region == "" || userPoolID == "" {
		panic("COGNITO_REGION and COGNITO_USER_POOL_ID must be set in environment variables")
	}
	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolID)
	ctx := context.Background()
	var err error
	kf, err = keyfunc.NewDefaultCtx(ctx, []string{jwksURL})
	if err != nil {
		panic(fmt.Sprintf("failed to create keyfunc.Keyfunc: %v", err))
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenString, kf.Keyfunc)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid JWT claims"})
			return
		}

		sub, ok := claims["sub"].(string)
		if !ok || sub == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "sub not found in token"})
			return
		}

		c.Set(ContextKeyUserID, sub)
		c.Next()
	}
}
