package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app *fiber.App, enforcer *casbin.Enforcer) {
	// Create a main API group under the "/api" path to organize all API endpoints
	apiGroup := app.Group("/api")
	// Register admin routes
	AdminRouter(apiGroup, enforcer)
	// Register auth routes
	AuthRouter(apiGroup)
	// Register user routes
	UserRouter(apiGroup)
	// Register player routes
	PlayRouter(apiGroup)
}
