package routes

import (
	"github.com/gofiber/fiber/v2"
)

func StartRoutes(app *fiber.App) {
	publicRoutes(app)
	protectedRoutes(app)
}
