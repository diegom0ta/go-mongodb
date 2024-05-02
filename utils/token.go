package utils

import (
	"errors"
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// GenerateToken creates a new JWT token
func GenerateToken(email string) (string, error) {
	if email == "" {
		return "", errors.New("email must be non-empty strings")
	}

	config, err := ParseYaml()
	if err != nil {
		log.Fatalf("Error parsing yaml: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email": email,
			"exp":   time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString([]byte(config.Secret.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
