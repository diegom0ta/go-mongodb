package services

import (
	"context"
	"strconv"

	db "github.com/diegom0ta/go-mongodb/internal/database"
	"github.com/diegom0ta/go-mongodb/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const pageSize = 10

// GetUsers returns a list of users
func GetUsers(c *fiber.Ctx) error {
	var users []models.User

	ctx := context.Background()

	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"]

	page := c.Params("page")
	pageInt, _ := strconv.Atoi(page)

	database := db.Client.Database("mydatabase")
	collection := database.Collection("users")

	filter := bson.M{"_id": bson.M{"$gt": userId}}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	offset := int64((pageInt - 1) * pageSize)

	options := options.Find().SetSkip(offset).SetLimit(int64(pageSize))

	cursor, err := collection.Find(ctx, nil, options)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &users); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"users": users,
		"count": count,
	})
}
