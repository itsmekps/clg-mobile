package main

import (
	"fiber-boilerplate/app/database"
	"fiber-boilerplate/app/logger"
	"fiber-boilerplate/app/middleware"
	"fiber-boilerplate/app/repository"
	"fiber-boilerplate/app/router"
	"fiber-boilerplate/app/service"
	"fiber-boilerplate/config"
	"fiber-boilerplate/internal/bootstrap"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	// Configuration parameters
	configPath := "."    // Path to your configuration directory
	configFile := ".env" // File name without extension
	configType := "env"  // File type: "yaml", "json", or "env"

	// Initialize configuration
	config.InitConfig(configPath, configFile, configType)

	// Initialize zap logger
	logger.InitLogger(config.GetConfig().AppEnv)
	// Access the zap logger
	zapLogger := logger.Log
	// Flushes buffer, if any
	defer logger.Log.Sync()

	// Initialize the database connection (e.g., MySQL, MongoDB)
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err) // Exit the application if the database initialization fails
	}

	// Initialize repositories (data access layer) with the database connection
	repos := repository.InitRepositories(db)

	// Initialize services (business logic layer) with the repositories
	service.InitServices(repos)

	// Initialize casbin - access control middleware
	enforcer, err := database.InitCasbinEnforcer()
	if err != nil {
		panic(err)
	}

	// Set up the Fiber web server
	app := bootstrap.InitWebServer()

	// Enable Cross-Origin Resource Sharing (CORS) with the specified configuration
	app.Use(cors.New(config.CORSConfig()))

	// Apply Request ID middleware
	app.Use(middleware.RequestIDMiddleware())

	// Apply Zap request logger middleware
	app.Use(middleware.ZapRequestLogger(zapLogger))

	// Apply Auth middleware globally
	// app.Use(middleware.AuthMiddleware())

	// Apply Casbin middleware globally
	// app.Use(middleware.CasbinMiddleware(enforcer))

	// Register a logging middleware to log incoming requests
	// app.Use(middleware.LogMiddleware())

	// Register any global or additional middleware
	middleware.RegisterMiddleware(app)

	// Register all API routes
	router.ApiRouter(app, enforcer)

	// Start the web server
	startServer(app)
}

func startServer(app *fiber.App) {

	log.Fatal(app.Listen(":" + config.GetConfig().AppPort))

	// Note: Add a deferred cleanup step here if a global database connection needs to be closed
	// defer db.MySQL.Close()
}
