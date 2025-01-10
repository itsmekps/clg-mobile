package handlers

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) AuthHandler {
	return AuthHandler{AuthService: authService}
}

// Login returns a access token by username and password
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	// Retrieve validated email and password from request context
	validatedRequest := c.Locals("validatedRequest").(*dtos.LoginRequest)

	// Extract email and password from the validated request
	email := validatedRequest.Email
	password := validatedRequest.Password

	// Call AuthService to handle login
	token, err := h.AuthService.Login(email, password)
	if err != nil {
		return err.Respond(c)
	}

	// Return success response
	return c.JSON(fiber.Map{
		"access_token": token,
		"token_type":   "Bearer",
	})
}
