package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	defaultExpireDuration = time.Hour * 8
)

// Generate generates a JSON Web Token with an 8-hour expiration
func Generate(secretKey string, data interface{}) (string, error) {
	// Define the secret key used to sign the token
	secretByte := []byte(secretKey)

	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(defaultExpireDuration).Unix() // Expiration time (8 hours from now)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretByte)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
