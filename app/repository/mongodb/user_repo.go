package mongodb

import (
	"context"
	"fiber-boilerplate/app/models"
	"fmt"

	// "log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}
func (r *UserRepository) GetUser(id primitive.ObjectID) (*models.User, error) {
	fmt.Println("objID repo: ", id)

	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.collection.InsertOne(context.TODO(), user)
	return err
}
func (r *UserRepository) UpdateUser(id int, user *models.User) error {
	_, err := r.collection.UpdateOne(context.TODO(), bson.M{"id": id}, bson.M{"$set": user})
	return err
}
func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.collection.DeleteOne(context.TODO(), bson.M{"id": id})
	return err
}
