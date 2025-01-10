// service/init.go
package service

import (
	"fiber-boilerplate/app/repository/mongodb"
	"log"
)

var UserServiceInstance *UserService
var PlayerServiceInstance *PlayerService
var AuthServiceInstance *AuthService

func InitServices(repos map[string]interface{}) {
	// Initialize MySQL services
	// mysqlRepos, ok := repos["mysql"].(map[string]interface{})
	// if !ok {
	// 	log.Fatal("MySQL repositories not found or invalid type")
	// }

	mongoRepos, ok := repos["mongodb"].(map[string]interface{})
	if !ok {
		log.Fatal("MySQL repositories not found or invalid type")
	}

	userRepo, ok := mongoRepos["userRepo"].(*mongodb.UserRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}
	UserServiceInstance = NewUserService(userRepo)
	AuthServiceInstance = NewAuthService(userRepo)

	playerRepo, ok := mongoRepos["playerRepo"].(*mongodb.PlayerRepository)
	if !ok {
		log.Fatal("Invalid user repository instance")
	}
	PlayerServiceInstance = NewPlayerService(playerRepo)

}
