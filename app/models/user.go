package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"` // MongoDB ObjectID
	Email         string             `bson:"email" json:"email"`      // User's email
	Password      string             `bson:"password" json:"password"` // Hashed password
	Avatar        string             `bson:"avatar,omitempty" json:"avatar,omitempty"` // Optional avatar URL
	Color         string             `bson:"color,omitempty" json:"color,omitempty"`   // Optional color
	FirstName     string             `bson:"first_name" json:"first_name"`             // User's first name
	LastName      string             `bson:"last_name" json:"last_name"`               // User's last name
	TwoStepsAuth  bool               `bson:"two_steps_auth" json:"two_steps_auth"`     // Two-factor authentication status
	Role          string             `bson:"role" json:"role"`                         // User's role (e.g., "user")
	Status        string             `bson:"status" json:"status"`                     // Account status (e.g., "active")
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`             // Account creation timestamp
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`             // Last update timestamp
}
