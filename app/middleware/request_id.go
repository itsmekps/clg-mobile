package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestIDMiddleware generates a unique UUID for each request
func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Generate a new UUID
		requestID := uuid.New().String()

		// Store the UUID in the context
		c.Locals("request_id", requestID)

		// Set the UUID in the response headers for client-side traceability
		c.Set("X-Request-ID", requestID)

		return c.Next()
	}
}
