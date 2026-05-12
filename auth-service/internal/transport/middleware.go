package transport

import (
	"net/http"
	"strings"

	"github.com/SussyaPusya/swiftTalk/auth-service/internal/pkg"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	j *pkg.JWTManager
}

func NewMiddleware(m *pkg.JWTManager) *Middleware {
	return &Middleware{
		j: m,
	}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := m.j.VerifyAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
