package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thoriqulumar/cats-social-service-w1/internal/pkg/jwt"
)

func AuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Perform your authorization logic here, e.g., checking JWT token, session, etc.
		// For this example, let's just check for a specific header.
		token := c.GetHeader("Authorization")

		fmt.Println("token", token)
		fmt.Println("JWTSecret", secretKey)

		claims, err := jwt.ValidateToken(token, secretKey)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		data := claims["data"]
		fmt.Println("Data from token:", data)

		// If authorization is successful, proceed to the next handler
		c.Next()
	}
}
