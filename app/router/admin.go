package router

import (
	"fiber-boilerplate/app/handlers/admin"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

// AdminRouter registers routes for managing Casbin policies dynamically
func AdminRouter(router fiber.Router, enforcer *casbin.Enforcer) {

	// Create a sub-group of routes under the "/admin" path
	adminGroup := router.Group("/admin")

	// Add and remove policies
	adminGroup.Post("/add-policy", admin.AddPolicyHandler(enforcer))
	adminGroup.Post("/remove-policy", admin.RemovePolicyHandler(enforcer))
	adminGroup.Get("/get-policies", admin.ListPoliciesHandler(enforcer))
}
