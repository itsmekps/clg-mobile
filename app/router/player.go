package router

import (
	"fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
)

func PlayRouter(router fiber.Router) {

	// Initialize a new instance of the validator
	// validate := validator.New()

	playerHandler := handlers.NewPlayerHandler((*service.PlayerService)(service.PlayerServiceInstance))

	// Create a sub-group of routes under the "/players" path
	playerGroup := router.Group("/players")

	{
		// get list of all players based on filters/query params
		playerGroup.Get("/", playerHandler.GetPlayersList)

		// get details of a player by playerID
		// playerGroup.Get("/:id", playerHandler.PlayerDetails)

		// search player by name(string pattern match)
		playerGroup.Get("/search", playerHandler.SearchPlayers)
	}
}
