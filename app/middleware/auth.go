package middleware

import (
	"fiber-boilerplate/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Exclude specific routes from authentication
		if strings.HasPrefix(c.Path(), "/api/auth/login") {
			return c.Next() // Skip middleware for login
		}

		// Get Authorization header
		authHeader := c.Get("Authorization")

		// Check if the Authorization header exists and starts with "Bearer "
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized: Missing or invalid token",
			})
		}

		// Extract the token from the Authorization header
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the JWT token
		isValid, claims, err := utils.ValidateTokenWithClaims(token)
		if err != nil || !isValid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized: Invalid or expired token",
			})
		}

		// Extract user_id from claims and convert it to ObjectID
		userIDStr := claims["user_id"].(string)
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   "Unauthorized: Invalid user ID",
			})
		}
		role := claims["role"].(string)
		// Store user_id as ObjectID in the context
		c.Locals("user_id", userID)
		// Set the role in Locals
		c.Locals("role", role)

		// Proceed to the next handler
		return c.Next()
	}
}
