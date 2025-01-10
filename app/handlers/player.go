package handlers

import (
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
)

type PlayerHandler struct {
	PlayerService *service.PlayerService
}

func NewPlayerHandler(playerService *service.PlayerService) PlayerHandler {
	return PlayerHandler{PlayerService: playerService}
}

// GetUser returns a user by ID
func (h *PlayerHandler) GetPlayersList(c *fiber.Ctx) error {

	// Default values
	page := 1
	limit := 20

	// Fetch players from the database
	// The service layer is responsible for interacting with MongoDB
	players, err := service.PlayerServiceInstance.GetPlayersList(page, limit)
	if err != nil {
		// Return the error if the user retrieval fails (e.g., user not found or DB error)
		return err.Respond(c)
	}

	// Send the JSON response back to the client, wrapping the user data in a success object
	return c.JSON(fiber.Map{
		"success": true,                          // Indicate the operation was successful
		"data":    fiber.Map{"players": players}, // Include the sanitized user data
	})
}

// func GetPlayerDetails(c *fiber.Ctx) error {
// 	// Retrieve user_id from the context
// 	userID := c.Locals("user_id").(primitive.ObjectID)

// 	// Fetch the user from the database using the validated ObjectID
// 	// The service layer is responsible for interacting with MongoDB
// 	user, err := service.UserServiceInstance.GetUser(userID)
// 	if err != nil {
// 		// Return the error if the user retrieval fails (e.g., user not found or DB error)
// 		return err
// 	}

// 	// Create a sanitized response object excluding sensitive fields (e.g., password)
// 	response := dtos.UserResponse{
// 		ID:        user.ID.Hex(),  // Convert the ObjectID to its string representation (Hex format)
// 		FirstName: user.FirstName, // User's first anme
// 		LastName:  user.LastName,  // User's last anme
// 		Email:     user.Email,     // User's email address
// 	}

// 	// Send the JSON response back to the client, wrapping the user data in a success object
// 	return c.JSON(fiber.Map{
// 		"success": true,                        // Indicate the operation was successful
// 		"data":    fiber.Map{"user": response}, // Include the sanitized user data
// 	})
// }
