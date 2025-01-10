package errors

import (
	"github.com/gofiber/fiber/v2"
)

// Predefined application errors
var (
	USER_NOT_FOUND = &AppError{
		StatusCode: fiber.StatusNotFound,
		Code:       "USER_NOT_FOUND",
		Msg:        "No such user found",
	}

	INVALID_PASSWORD = &AppError{
		StatusCode: fiber.StatusUnauthorized,
		Code:       "INVALID_PASSWORD",
		Msg:        "The password you provided is incorrect",
	}

	INTERNAL_SERVER_ERROR = &AppError{
		StatusCode: fiber.StatusInternalServerError,
		Code:       "INTERNAL_SERVER_ERROR",
		Msg:        "An internal server error occurred",
	}
)

// AppError struct for error responses
type AppError struct {
	StatusCode int    // HTTP status code
	Code       string // Error code for identifying the error
	Msg        string // Error message for display
}

// Respond sends the error response in the desired JSON format
func (e *AppError) Respond(c *fiber.Ctx) error {
	return c.Status(e.StatusCode).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code": e.Code,
			"msg":  e.Msg,
		},
	})
}
