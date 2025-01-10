package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Position struct {
	Resource string  `bson:"resource" json:"resource"` // Resource type (e.g., "positions")
	ID       float64 `bson:"id" json:"id"`             // Position ID
	Name     string  `bson:"name" json:"name"`         // Position name (e.g., "Batsman")
}

type Status struct {
	T20  bool `bson:"T20" json:"T20"`
	ODI  bool `bson:"ODI" json:"ODI"`
	TEST bool `bson:"TEST" json:"TEST"`
}

type Player struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`                               // MongoDB ObjectID
	UUID         string             `bson:"uuid" json:"uuid"`                                      // Unique identifier
	Position     Position           `bson:"position" json:"position"`                              // Player's position
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`                          // Last update timestamp
	PlayerID     float64            `bson:"id" json:"player_id"`                                   // Player ID
	BattingStyle string             `bson:"battingstyle" json:"batting_style"`                     // Batting style (e.g., "right-hand-bat")
	BowlingStyle string             `bson:"bowlingstyle,omitempty" json:"bowling_style,omitempty"` // Bowling style
	ImagePath    string             `bson:"image_path" json:"image_path"`                          // Image URL
	DateOfBirth  string             `bson:"dateofbirth" json:"date_of_birth"`                      // Date of birth (YYYY-MM-DD)
	CountryID    float64            `bson:"country_id" json:"country_id"`                          // Country ID
	Gender       string             `bson:"gender" json:"gender"`                                  // Gender (e.g., "m")
	FirstName    string             `bson:"firstname" json:"first_name"`                           // First name
	LastName     string             `bson:"lastname" json:"last_name"`                             // Last name
	FullName     string             `bson:"fullname" json:"full_name"`                             // Full name
	Status       Status             `bson:"status" json:"status"`                                  // Status for formats (T20, ODI, TEST)
}

type PlayerList struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`                // MongoDB ObjectID
	Avatar    string             `bson:"avatar" json:"avatar"`         // Avatar URL
	Name      string             `bson:"name" json:"name"`             // Player name
	Country   float64            `bson:"country" json:"country"`       // Country ID
	Role      string             `bson:"role" json:"role"`             // Player role
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"` // Last updated timestamp
	Status    Status             `bson:"status" json:"status"`
}
