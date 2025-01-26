package mongodb

import (
	"context"
	"fiber-boilerplate/app/models"

	// "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PlayerRepository struct {
	collection *mongo.Collection
}

func NewPlayerRepository(collection *mongo.Collection) *PlayerRepository {
	return &PlayerRepository{collection: collection}
}

func (r *PlayerRepository) PlayerList(page, limit int) ([]models.PlayerList, int64, error) {
	var players []models.PlayerList

	// Get the total count of documents in the collection
	totalCount, err := r.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, 0, err
	}

	// Pagination: Skip (page-1)*limit and limit to `limit` results
	options := options.Find().
		SetLimit(int64(limit)).
		SetSkip(int64((page - 1) * limit))

	// Fetch data with pagination
	cursor, err := r.collection.Find(context.TODO(), bson.M{}, options)
	if err != nil {
		return nil, 0, err
	}

	if err = cursor.All(context.TODO(), &players); err != nil {
		return nil, 0, err
	}

	return players, totalCount, nil
}

func (r *PlayerRepository) SearchPlayers(query string, limit int) ([]models.PlayerSearch, error) {
	var players []models.PlayerSearch

	// Create a MongoDB query to search for players by firstname
	filter := bson.M{
		"firstname": bson.M{
			"$regex":   query,
			"$options": "i", // Case-insensitive search
		},
	}

	// Pagination: Skip (page-1)*limit and limit to `limit` results
	options := options.Find().
		SetLimit(int64(limit))

	// Fetch data with pagination
	cursor, err := r.collection.Find(context.TODO(), filter, options)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.TODO(), &players); err != nil {
		return nil, err
	}

	return players, nil
}

func (r *PlayerRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PlayerRepository) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}
func (r *PlayerRepository) UpdateUser(id int, user *models.User) error {
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": user})
	return err
}
func (r *PlayerRepository) DeleteUser(id int) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}
