package server

import (
	"fmt"
	"log"

	"github.com/diegom0ta/go-mongodb/internal/http/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App

func Run(port uint) {
	app = fiber.New()

	app.Use(cors.New())

	routes.StartRoutes(app)

	if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		log.Printf("Error listening server: %v", err)
	}
}

func Shutdown() {
	err := app.Shutdown()
	if err != nil {
		log.Printf("Error while shutting down: %v", err)
	}

	log.Println("Server closed")
}
