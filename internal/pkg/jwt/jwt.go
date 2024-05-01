package jwt

import (
	"fmt"
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

func ValidateToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key for validation
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("err parse token", err)
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("unable to parse claims")
}
