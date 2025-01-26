package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	LastName     string             `json:"lastname" bson:"lastname"`
	FullName     string             `json:"fullname" bson:"fullname"`
	Image        Image              `json:"image" bson:"image"`
	SmPlayersID  float64            `json:"sm_players_id" bson:"sm_players_id"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	Gender       string             `json:"gender" bson:"gender"`
	BattingStyle string             `json:"battingstyle" bson:"battingstyle"`
	BowlingStyle string             `json:"bowlingstyle" bson:"bowlingstyle"`
	Position     Position           `json:"position" bson:"position"`
	UpdatedAt    string             `json:"updated_at" bson:"updated_at"`
	SmCountryID  float64            `json:"_sm_country_id" bson:"_sm_country_id"`
	DateOfBirth  string             `json:"dateofbirth" bson:"dateofbirth"`
	CountriesID  primitive.ObjectID `json:"_countries_id" bson:"_countries_id"`
}

type Image struct {
	Host     string `json:"host" bson:"host"`
	Path     string `json:"path" bson:"path"`
	Filename string `json:"filename" bson:"filename"`
}

type Position struct {
	SmPositionsID float64 `json:"_sm_positions_id" bson:"_sm_positions_id"`
	Name          string  `json:"name" bson:"name"`
}

// for playerlist api - /players/
type PlayerList struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	LastName     string             `json:"ln" bson:"lastname"`
	FullName     string             `json:"fun" bson:"fullname"`
	Image        Image              `json:"img" bson:"image"`
	FirstName    string             `json:"fn" bson:"firstname"`
	Gender       string             `json:"sex" bson:"gender"`
	BattingStyle string             `json:"bts" bson:"battingstyle"`
	BowlingStyle string             `json:"bws" bson:"bowlingstyle"`
	Position     Position           `json:"pos" bson:"position"`
	DateOfBirth  string             `json:"dob" bson:"dateofbirth"`
}

// for player search api - /players/search
type PlayerSearch struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	FullName string             `json:"fun" bson:"fullname"`
	Image    Image              `json:"img" bson:"image"`
}
