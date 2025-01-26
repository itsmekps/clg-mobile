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

// func (r *PlayerRepository) FindAll() ([]models.Player, error) {
//     var players []models.Player

//     // Fetch all players from the MongoDB collection
//     cursor, err := r.collection.Find(context.TODO(), bson.M{})
//     if err != nil {
//         return nil, err
//     }

//     // Decode the results into the Player slice
//     if err = cursor.All(context.TODO(), &players); err != nil {
//         return nil, err
//     }

//     return players, nil
// }

func (r *PlayerRepository) FindAll(fields bson.M, page, limit int) ([]models.Player, error) {
    var players []models.Player

    // Validate page number
    if page < 1 {
        page = 1
    }

    // Validate limit
    if limit < 1 {
        limit = 10 // Default limit
    }

    // Pagination: Skip (page-1)*limit and limit to `limit` results
    options := options.Find().
        SetLimit(int64(limit)).
        SetSkip(int64((page - 1) * limit))

    // Add projection if fields are specified
    if fields != nil {
        options.SetProjection(fields)
    }

    // Fetch data with pagination and projection
    cursor, err := r.collection.Find(context.TODO(), bson.M{}, options)
    if err != nil {
        return nil, err
    }

    if err = cursor.All(context.TODO(), &players); err != nil {
        return nil, err
    }

    return players, nil
}

// func (r *PlayerRepository) GetPlayersList(page, limit int) ([]models.PlayerList, error) {
// 	if page <= 0 {
// 		page = 1
// 	}
// 	if limit <= 0 {
// 		limit = 20
// 	}

// 	pipeline := mongo.Pipeline{
// 		bson.D{
// 			{Key: "$project", Value: bson.D{
// 				{Key: "avatar", Value: "$image_path"},
// 				{Key: "name", Value: "$fullname"},
// 				{Key: "country", Value: "$country_id"},
// 				{Key: "role", Value: "$position.name"},
// 				{Key: "updated_at", Value: "$updated_at"},
// 			}},
// 		},
// 		bson.D{
// 			{Key: "$sort", Value: bson.D{
// 				{Key: "updated_at", Value: -1},
// 			}},
// 		},
// 		bson.D{
// 			{Key: "$skip", Value: int64((page - 1) * limit)},
// 		},
// 		bson.D{
// 			{Key: "$limit", Value: int64(limit)},
// 		},
// 	}

// 	cursor, err := r.collection.Aggregate(context.TODO(), pipeline)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to execute aggregation pipeline: %w", err)
// 	}
// 	defer cursor.Close(context.TODO())

// 	// Decode directly into the Player struct
// 	var players []models.PlayerList
// 	if err := cursor.All(context.TODO(), &players); err != nil {
// 		return nil, fmt.Errorf("failed to decode results: %w", err)
// 	}

// 	return players, nil
// }

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
