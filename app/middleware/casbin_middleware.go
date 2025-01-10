package middleware

import (
	"fmt"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
)

func CasbinMiddleware(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Exclude specific routes from authentication
		if strings.HasPrefix(c.Path(), "/api/auth/login") {
			return c.Next() // Skip middleware for login
		}
		// Extract role from JWT claims
		role, ok := c.Locals("role").(string)
		if !ok || role == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized: Missing or invalid role",
			})
		}

		// Extract object (path) and action (method)
		obj := c.Path()   // Resource path
		act := c.Method() // HTTP method

		// Debug: Print enforcement inputs
		fmt.Printf("Enforcing with: sub=%s, obj=%s, act=%s\n", role, obj, act)

		// Enforce the policy
		allowed, err := e.Enforce(role, obj, act)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "Error checking permissions: " + err.Error(),
			})
		}

		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"success": false,
				"error":   "Forbidden: Access denied",
			})
		}

		return c.Next()
	}
}
