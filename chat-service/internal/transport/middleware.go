package transport

import (
	"log"
	"net/http"
	"strings"

	pb "github.com/SussyaPusya/swiftTalk/chat-service/client"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	client pb.Account_ServiceClient
}

func NewMiddleware(cl pb.Account_ServiceClient) *Middleware {
	return &Middleware{client: cl}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		resp, err := m.client.ValidateToken(c, &pb.ValidateTokenRequest{Token: tokenStr})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		if !resp.IsValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("user_id", resp.UserId)
		log.Println("user_id:", resp.UserId)
	}
}
