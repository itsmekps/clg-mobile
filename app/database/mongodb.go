package database

import (
	"context"
	"fiber-boilerplate/config"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient is the global MongoDB client instance that can be used across the application.
var MongoClient *mongo.Client

// InitMongoDB initializes the MongoDB connection.
// It loads configuration values, creates a connection client, and verifies
// the connection by pinging the MongoDB server. If successful, it assigns
// the client to the global MongoClient variable and returns it.
func InitMongoDB() *mongo.Client {

	// Load configuration values (e.g., MongoDB user, password, host) from environment or config file.
	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err) // Exit the application if the configuration fails to load
	}

	// Create a context with a 10-second timeout for connection operations.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Build the MongoDB connection string and create a new client.
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority&appName=CLG",
			v.GetString("Mongodb_user"),
			v.GetString("Mongodb_password"),
			v.GetString("Mongodb_host"),
		)),
	)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Verify the connection by pinging the MongoDB server.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")

	// Assign the client to the global MongoClient variable so it can be reused.
	MongoClient = client
	return client
}

// GetCollection returns a reference to a specific MongoDB collection.
// It accepts a database name and a collection name, and returns the corresponding collection object.
// If the global MongoClient is nil, it logs a fatal error indicating the client is not initialized.
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}

// DisconnectMongoDB closes the MongoDB connection if the global MongoClient is initialized.
// It creates a new context with a 10-second timeout, attempts to disconnect from MongoDB,
// and logs the result. If disconnection fails, it logs a fatal error.
func DisconnectMongoDB() {
	if MongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := MongoClient.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
		log.Println("Successfully disconnected from MongoDB")
	}
}
