package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func ApiRouterV1(app *fiber.App, enforcer *casbin.Enforcer) {
	// Create a main API group under the "/api" path to organize all API endpoints
	apiV1 := app.Group("/api/v1")
	// Register admin routes
	AdminRouter(apiV1, enforcer)
	// Register auth routes
	AuthRouter(apiV1)
	// Register user routes
	UserRouter(apiV1)
	// Register player routes
	PlayRouter(apiV1)
	// Register news routes
	NewsRouterV1(apiV1)
}

// for future use - api versioning - v2

// func ApiRouterV2(app *fiber.App, enforcer *casbin.Enforcer) {
// 	// Create a main API group under the "/api" path to organize all API endpoints
// 	apiV2 := app.Group("/api/v2")

// 	NewsRouterV2(apiV2)
// }
