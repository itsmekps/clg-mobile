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

var MongoClient *mongo.Client

// ConnectMongoDB initializes the MongoDB connection.
func InitMongoDB() *mongo.Client {

	v, err := config.InitConfig()
	if err != nil {
		log.Fatal(err) // Exit the application if the configuration fails to load
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Successfully connected to MongoDB")
	MongoClient = client
	return client
}

// GetCollection returns a MongoDB collection.
func GetCollection(databaseName, collectionName string) *mongo.Collection {
	if MongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return MongoClient.Database(databaseName).Collection(collectionName)
}

// DisconnectMongoDB closes the MongoDB connection.
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
