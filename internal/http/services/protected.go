package services

import (
	"context"
	"fmt"

	db "github.com/diegom0ta/go-mongodb/internal/database"
	"github.com/diegom0ta/go-mongodb/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUsers returns a list of users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	ctx := context.Background()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"]

	database := db.Client.Database("mydatabase")
	collection := database.Collection("users")

	filter := bson.M{"email": email}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return fmt.Errorf("failed to retrieve users: %v", err)
	}
	defer cur.Close(ctx)

	if err := cur.All(ctx, &users); err != nil {
		return fmt.Errorf("failed to list users: %v", err)
	}

	return c.JSON(users)
}
