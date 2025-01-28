// Package service provides the initialization of services that use repositories
// to access data. This file specifically handles the creation of service
// instances by injecting the appropriate repositories into them.
package service

import (
	"fiber-boilerplate/app/repository/mongodb"
	"log"
)

// Global variables holding instances of each service for application-wide usage.
var (
	UserServiceInstance   *UserService
	PlayerServiceInstance *PlayerService
	AuthServiceInstance   *AuthService
	NewsServiceInstance   *NewsService
)

// InitServices initializes all services required by the application.
// It expects a map of repositories (currently handling "mongodb") to be passed in.
// Each repository is type-asserted to the appropriate repository struct before
// being used to create the corresponding service instances.
//
// If any repository is not found or has the wrong type, the application will log
// a fatal error and exit.
func InitServices(repos map[string]interface{}) {

	// Initialize MySQL services
	// mysqlRepos, ok := repos["mysql"].(map[string]interface{})
	// if !ok {
	// 	log.Fatal("MySQL repositories not found or invalid type")
	// }

	// Attempt to retrieve the MongoDB repositories from the input map
	mongoRepos, ok := repos["mongodb"].(map[string]interface{})
	if !ok {
		log.Fatal("Mongo repositories not found or invalid type")
	}

	// Retrieve and validate the user repository
	userRepo, ok := mongoRepos["userRepo"].(*mongodb.UserRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}

	// Create the user-related services
	UserServiceInstance = NewUserService(userRepo)
	AuthServiceInstance = NewAuthService(userRepo)

	// Retrieve and validate the player repository
	playerRepo, ok := mongoRepos["playerRepo"].(*mongodb.PlayerRepository)
	if !ok {
		log.Fatal("Invalid player repository instance")
	}

	// Create the player service
	PlayerServiceInstance = NewPlayerService(playerRepo)

	// Retrieve and validate the news repository
	newsRepo, ok := mongoRepos["newsRepo"].(*mongodb.NewsRepository)
	if !ok {
		log.Fatal("Invalid news repository instance")
	}

	// Create the news service
	NewsServiceInstance = NewNewsService(newsRepo)
}
