package middleware

import (
	"fmt"
	"log"
	"strings"

	"github.com/diegom0ta/go-mongodb/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtAuth() fiber.Handler {
	config, err := utils.ParseYaml()
	if err != nil {
		log.Fatalf("Error parsing yaml: %v", err)
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.SendStatus(fiber.StatusForbidden)
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.SendStatus(fiber.StatusForbidden)
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(config.Secret.SecretKey), nil
		})

		if err != nil || !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		c.Locals("user", token)
		return c.Next()
	}
}
