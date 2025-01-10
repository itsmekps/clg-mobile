package router

import (
	"fiber-boilerplate/app/dtos"
	"fiber-boilerplate/app/handlers"
	"fiber-boilerplate/app/messages"
	"fiber-boilerplate/app/middleware"
	"fiber-boilerplate/app/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(router fiber.Router) {
	validate := validator.New()

	// authHandler := handlers.NewAuthHandler((*service.AuthService)(service.UserServiceInstance))
	userHandler := handlers.NewUserHandler((*service.UserService)(service.UserServiceInstance))

	// Create a sub-group of routes under the "/users" path
	userGroup := router.Group("/users")

	{
		// Validate details request without parameters
		userGroup.Get("/details", userHandler.GetUserDetails)

		userGroup.Get("/:id", middleware.ValidateRequestDTO(validate, &dtos.GetUserByMongoID{}, messages.GetUser), userHandler.GetUser)
		// userGroup.Get("/", user.GetUser)
	}
}
