package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// InitRepositories initializes and returns a map of MongoDB-based repository instances.
// Each repository uses a specific collection from the provided MongoDB client.
// The returned map can be used by service layers to perform data operations.
//
// Parameters:
//   - client: a *mongo.Client representing the connection to MongoDB.
//
// Returns:
//   - A map[string]interface{} where each key corresponds to a repository name
//     (e.g., "userRepo", "playerRepo", etc.), and the value is the initialized repository instance.
//   - An error if any step in creating repositories fails (currently always returns nil).
//
// Usage:
//
//	repos, err := InitRepositories(mongoClient)
//	if err != nil {
//	    log.Fatal("Failed to initialize repositories: ", err)
//	}
//	userRepo := repos["userRepo"].(*UserRepository)
//	// ...
func InitRepositories(client *mongo.Client) (map[string]interface{}, error) {
	// Retrieve a reference to the desired database from the MongoDB client.
	db := client.Database("clg")

	// Create a map to hold references to each repository.
	// The repositories are instantiated with their respective MongoDB collections.
	repos := map[string]interface{}{
		"userRepo":   NewUserRepository(db.Collection("users")),
		"playerRepo": NewPlayerRepository(db.Collection("players")),
		"newsRepo":   NewNewsRepository(db.Collection("news")),
		// Add other repository instances here as needed.
	}

	// Return the map of repositories and a nil error (no errors handled at the moment).
	return repos, nil
}
