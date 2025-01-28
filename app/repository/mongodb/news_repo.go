package mongodb

import (
	"context"
	"time"

	"fiber-boilerplate/app/models"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsRepository struct {
	Collection *mongo.Collection
}

func NewNewsRepository(collection *mongo.Collection) *NewsRepository {
	return &NewsRepository{Collection: collection}
}

func (r *NewsRepository) CreateNews(news *models.News) error {
	news.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	news.UpdatedAt = news.CreatedAt
	_, err := r.Collection.InsertOne(context.TODO(), news)
	return err
}

func (r *NewsRepository) GetNewsByID(id primitive.ObjectID) (*models.News, error) {
	var news models.News
	err := r.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&news)
	return &news, err
}

func (r *NewsRepository) UpdateNews(id primitive.ObjectID, update bson.M) error {
	_, err := r.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *NewsRepository) DeleteNews(id primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

func (r *NewsRepository) ListNews() ([]models.News, error) {
	cursor, err := r.Collection.Find(context.TODO(), bson.M{})
	if err != nil {
		// return nil, err
		// Wrap the error with context
		return nil, errors.Wrap(err, "repository.ListNews: failed to execute find query")
	}
	defer cursor.Close(context.TODO())

	var newsList []models.News
	for cursor.Next(context.TODO()) {
		var news models.News
		if err := cursor.Decode(&news); err != nil {
			// return nil, err
			// Wrap the error with context
			return nil, errors.Wrap(err, "repository.ListNews: failed to decode news documents")
		}
		newsList = append(newsList, news)
	}
	return newsList, nil
}
