package routes

import (
	"github.com/diegom0ta/go-mongodb/internal/http/middleware"
	srv "github.com/diegom0ta/go-mongodb/internal/http/services"
	"github.com/gofiber/fiber/v2"
)

func protectedRoutes(app *fiber.App) fiber.Router {
	route := app.Group("/protected")

	route.Use(middleware.JwtAuth())
	route.Get("/users/:page", srv.GetUsers)

	return route
}
