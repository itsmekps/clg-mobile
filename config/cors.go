package config

import "github.com/gofiber/fiber/v2/middleware/cors"

// CORSConfig returns the configuration for CORS middleware
func CORSConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     "http://localhost:5173, https://example.com",
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Content-Type, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}
}
