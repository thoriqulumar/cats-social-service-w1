package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/jwt"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform your authorization logic here, e.g., checking JWT token, session, etc.
		// For this example, let's just check for a specific header.
		auth := c.GetHeader("Authorization")

		// Check if the Authorization header is empty
		if auth == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is missing"})
			c.Abort()
			return
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract the token from the Authorization header
		token := strings.TrimPrefix(auth, "Bearer ")

		claims, err := jwt.ValidateToken(token, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		data := claims["data"]
		// Add user data to the request context
		c.Set("userData", data)

		// If authorization is successful, proceed to the next handler
		c.Next()
	}
}
