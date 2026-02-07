package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/AnuragDahiwade/task-manager/internal/auth"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		// Expect: Bearer <token>
		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token := parts[1]

		claims, err := auth.VerifyToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user info in context
		uid, ok := claims["user_id"].(string)

		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token data"})
			c.Abort()
			return
		}

		c.Set("user_id", uid)

		role := fmt.Sprint(claims["role"])

		c.Set("role", role)

		c.Next()
	}
}
