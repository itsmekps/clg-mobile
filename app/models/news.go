package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// News represents the main news structure
type News struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title" json:"title"`
	Description string               `bson:"description" json:"description"`
	Content     string               `bson:"content" json:"content"`
	Assets      []primitive.ObjectID `bson:"assets" json:"assets"` // References to Asset IDs
	Author      string               `bson:"author" json:"author"`
	PublishedAt primitive.DateTime   `bson:"published_at" json:"published_at"`
	CreatedAt   primitive.DateTime   `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime   `bson:"updated_at" json:"updated_at"`
}
