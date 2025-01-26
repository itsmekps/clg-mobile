package handlers

import (
	"fiber-boilerplate/app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PlayerHandler struct {
	PlayerService *service.PlayerService
}

func NewPlayerHandler(playerService *service.PlayerService) PlayerHandler {
	return PlayerHandler{PlayerService: playerService}
}

// Get players list
func (h *PlayerHandler) GetPlayersList(c *fiber.Ctx) error {
	// Extract query parameters with default values
	page, _ := strconv.Atoi(c.Query("page", "1"))    // Default page: 1
	limit, _ := strconv.Atoi(c.Query("limit", "20")) // Default limit: 20

	// Validate page and limit
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 20 // Set a maximum limit to prevent excessive data retrieval
	}

	// Fetch players and pagination metadata from the database
	players, pagination, err := service.PlayerServiceInstance.GetPlayersList(page, limit)
	if err != nil {
		// Return the error if the player retrieval fails
		return err.Respond(c)
	}

	// Send the JSON response back to the client
	return c.JSON(fiber.Map{
		"success": true, // Indicate the operation was successful
		"data": fiber.Map{
			"players": players, // Include the player data
			"pagination": fiber.Map{
				"current_page": pagination.CurrentPage,
				"next_page":    pagination.NextPage,
				"prev_page":    pagination.PrevPage,
				"total_pages":  pagination.TotalPages,
			},
		},
	})
}

// Search player by matching name
func (h *PlayerHandler) SearchPlayers(c *fiber.Ctx) error {
	// Extract the search query from the request
	query := c.Query("q", "") // "q" is the query parameter for the search term

	// Validate the search query length
	if len(query) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Search query must be at least 3 characters long",
		})
	}

	// Extract pagination parameters with default values
	limit, _ := strconv.Atoi(c.Query("limit", "20")) // Default limit: 20

	// Validate limit
	if limit < 1 || limit > 100 {
		limit = 20 // Set a maximum limit to prevent excessive data retrieval
	}

	// Fetch players matching the search query
	players, err := service.PlayerServiceInstance.SearchPlayers(query, limit)
	if err != nil {
		// Return the error if the search fails
		return err.Respond(c)
	}

	// Send the JSON response back to the client
	return c.JSON(fiber.Map{
		"success": true, // Indicate the operation was successful
		"data": fiber.Map{
			"players": players, // Include the player data
		},
	})
}
