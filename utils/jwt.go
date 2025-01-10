package utils

import (
	"errors"
	"fiber-boilerplate/app/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

func GenerateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.Hex(),
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	fmt.Printf("Generated Claims: %+v\n", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates the given JWT and ensures it's valid
func ValidateTokenWithClaims(tokenString string) (bool, jwt.MapClaims, error) {
	// Parse the token with the secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	// Check for parsing errors or invalid tokens
	if err != nil || !token.Valid {
		return false, nil, err
	}

	// Extract claims and validate them
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, claims, nil
	}

	return false, nil, errors.New("invalid token claims")
}
