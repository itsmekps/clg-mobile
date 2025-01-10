package admin

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PolicyRequest struct {
	Sub string `json:"sub" validate:"required"` // Subject (role or user)
	Obj string `json:"obj" validate:"required"` // Object (resource path)
	Act string `json:"act" validate:"required"` // Action (HTTP method)
}

var validate = validator.New()

func AddPolicyHandler(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request PolicyRequest

		// Parse the JSON body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid JSON payload",
			})
		}

		// Validate the JSON payload
		if err := validate.Struct(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Validation failed",
				"details": err.Error(),
			})
		}

		// Add the policy dynamically
		if _, err := e.AddPolicy(request.Sub, request.Obj, request.Act); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "Failed to add policy: " + err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "Policy added successfully",
		})
	}
}

// RemovePolicyHandler handles removing policies dynamically
func RemovePolicyHandler(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request PolicyRequest

		// Parse the JSON body
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Invalid JSON payload",
			})
		}

		// Validate the JSON payload
		if err := validate.Struct(request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "Validation failed",
				"details": err.Error(),
			})
		}

		// Remove the policy dynamically
		if _, err := e.RemovePolicy(request.Sub, request.Obj, request.Act); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
				"error":   "Failed to remove policy: " + err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "Policy removed successfully",
		})
	}
}

// ListPoliciesHandler lists all Casbin policies
func ListPoliciesHandler(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		policies := e.GetPolicy()
		return c.JSON(fiber.Map{
			"success":  true,
			"policies": policies,
		})
	}
}
