package handlers

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) UserHandler {
	return UserHandler{UserService: userService}
}

// GetUser returns a user by ID
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	// Retrieve the validated request object from Fiber's context
	// The validated request contains the parsed and validated ObjectID from the route parameter
	validatedRequest := c.Locals("validatedRequest").(*dtos.GetUserByMongoID)

	// Fetch the user from the database using the validated ObjectID
	// The service layer is responsible for interacting with MongoDB
	user, err := service.UserServiceInstance.GetUser(validatedRequest.ID)
	if err != nil {
		// Return the error if the user retrieval fails (e.g., user not found or DB error)
		// return err
		return err.Respond(c)
	}

	// Create a sanitized response object excluding sensitive fields (e.g., password)
	response := dtos.UserResponse{
		ID:        user.ID.Hex(),  // Convert the ObjectID to its string representation (Hex format)
		FirstName: user.FirstName, // User's first anme
		LastName:  user.LastName,  // User's last anme
		Email:     user.Email,     // User's email address
	}

	// Send the JSON response back to the client, wrapping the user data in a success object
	return c.JSON(fiber.Map{
		"success": true,                        // Indicate the operation was successful
		"data":    fiber.Map{"user": response}, // Include the sanitized user data
	})
}

func (h *UserHandler) GetUserDetails(c *fiber.Ctx) error {
	// Retrieve user_id from the context
	userID := c.Locals("user_id").(primitive.ObjectID)

	// Fetch the user from the database using the validated ObjectID
	// The service layer is responsible for interacting with MongoDB
	user, err := service.UserServiceInstance.GetUser(userID)
	if err != nil {
		// Return the error if the user retrieval fails (e.g., user not found or DB error)
		return err.Respond(c)
	}

	// Create a sanitized response object excluding sensitive fields (e.g., password)
	response := dtos.UserResponse{
		ID:        user.ID.Hex(),  // Convert the ObjectID to its string representation (Hex format)
		FirstName: user.FirstName, // User's first anme
		LastName:  user.LastName,  // User's last anme
		Email:     user.Email,     // User's email address
	}

	// Send the JSON response back to the client, wrapping the user data in a success object
	return c.JSON(fiber.Map{
		"success": true,                        // Indicate the operation was successful
		"data":    fiber.Map{"user": response}, // Include the sanitized user data
	})
}
