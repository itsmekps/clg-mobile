package repository

import (
	"fiber-boilerplate/app/database"
	"fiber-boilerplate/app/repository/mongodb"
	// "fiber-boilerplate/app/repository/mysql"
	"log"
)

func InitRepositories(connections database.DBConnections) map[string]interface{} {

	// Initialize MongoDB repositories
	mongoRepos, err := mongodb.InitRepositories(connections.MongoDB)
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB repositories: %v", err)
	}

	// mysqlRepos := mysql.InitRepositories(connections.MySQL)

	repos := map[string]interface{}{
		// "mysql":   mysqlRepos,
		"mongodb": mongoRepos,
	}
	return repos
}
