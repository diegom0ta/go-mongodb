package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	db "github.com/diegom0ta/go-mongodb/internal/database"
	"github.com/diegom0ta/go-mongodb/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name                 string `json:"name"`
	Document             string `json:"document"`
	Email                string `json:"email"`
	Phone                string `json:"phone"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

func Register(c *fiber.Ctx) error {
	ctx := context.Background()

	newUser := new(User)

	if err := c.BodyParser(&newUser); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	database := db.Client.Database("mydatabase")
	collection := database.Collection("users")

	index := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, index)
	if err != nil {
		return fmt.Errorf("error creating unique index: %v", err)
	}

	uuid := uuid.New()

	pwdHash, err := validatePasswd(newUser.Password, newUser.PasswordConfirmation)
	if err != nil {
		return fmt.Errorf("password not valid: %v", err)
	}

	user := models.User{
		ID:        uuid.String(),
		Name:      newUser.Name,
		Document:  newUser.Document,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
		PwdHash:   pwdHash,
		CreatedAt: time.Now().UTC(),
	}

	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func hashPasswd(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %v", err)
	}

	return string(hashedPassword), nil
}

func validatePasswd(password, passwdConfirmation string) (string, error) {
	var hashedPasswd string

	if password == passwdConfirmation {
		hash, err := hashPasswd(password)
		if err != nil {
			return "", err
		}

		hashedPasswd = hash
	} else {
		return "", errors.New("passwords not equal")
	}

	return hashedPasswd, nil
}
