package router

import (
	"fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/service"

	"github.com/gofiber/fiber/v2"
)

// NewsRouter sets up the routes for news-related operations, such as creating, listing,
// retrieving by ID, updating, and deleting news items. It uses the NewsHandler
// to handle all the incoming requests and responses for /news.
func NewsRouter(router fiber.Router) {

	// Create a new NewsHandler instance by injecting the NewsService instance.
	newsHandler := handlers.NewNewsHandler((*service.NewsService)(service.NewsServiceInstance))

	// Create a sub-group of routes under the "/news" path.
	newsGroup := router.Group("/news")

	// Define all routes related to news items within this group.
	{
		// Route to create a new news item.
		newsGroup.Post("/", newsHandler.CreateNews)

		// Route to list all news items.
		newsGroup.Get("/", newsHandler.ListNews)

		// Route to retrieve a specific news item by its ID.
		newsGroup.Get("/:id", newsHandler.GetNewsByID)

		// Route to update an existing news item by its ID.
		newsGroup.Put("/:id", newsHandler.UpdateNews)

		// Route to delete a specific news item by its ID.
		newsGroup.Delete("/:id", newsHandler.DeleteNews)
	}
}
