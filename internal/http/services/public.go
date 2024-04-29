package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	db "github.com/diegom0ta/go-mongodb/internal/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
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

	filter := `{"email": "` + newUser.Email + `"}`

	var filterMap map[string]interface{}

	err := json.Unmarshal([]byte(filter), &filterMap)
	if err != nil {
		return fmt.Errorf("error decoding JSON filter: %v", err)
	}

	database := db.Client.Database("mydatabase")
	collection := database.Collection("users")

	// Find one document that matches the filter
	var user User
	err = collection.FindOne(context.Background(), filterMap).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("user with email '%s' does not exist", newUser.Email)
		} else {
			log.Printf("error checking user existence: %v", err)
		}
	}

	pwdHash, err := validatePasswd(newUser.Password, newUser.PasswordConfirmation)
	if err != nil {
		log.Printf("password not valid: %v", err)
	}

	_, err = collection.InsertOne(ctx, map[string]interface{}{
		"name":         newUser.Name,
		"document":     newUser.Document,
		"email":        newUser.Email,
		"phone":        newUser.Phone,
		"passwordHash": pwdHash,
	})
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
	}

	return c.SendStatus(200)
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
