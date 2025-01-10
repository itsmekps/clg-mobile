package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterMiddleware(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Protect API routes with auth middleware
	// api := app.Group("/api")
	// api.Use(AuthMiddleware())
}
