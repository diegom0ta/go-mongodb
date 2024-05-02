package routes

import (
	"log"

	srv "github.com/diegom0ta/go-mongodb/internal/http/services"
	"github.com/gofiber/fiber/v2"
)

func publicRoutes(app *fiber.App) fiber.Router {
	route := app.Group("/")

	route.Get("/", func(c *fiber.Ctx) error {
		log.Println("Main route")
		return c.SendString("Hello from diegom0ta")
	})
	route.Get("/routes", func(c *fiber.Ctx) error { return c.JSON(app.Stack()) })
	route.Post("/register", srv.Register)
	route.Post("/login/protected", srv.Login)

	return route
}
