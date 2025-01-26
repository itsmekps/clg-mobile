package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type League struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	UpdatedAt   string             `json:"updated_at" bson:"updated_at"`
	SmLeaguesID int                `json:"sm_leagues_id" bson:"sm_leagues_id"`
	SmSeasonID  int                `json:"_sm_season_id" bson:"_sm_season_id"`
	SmCountryID int                `json:"_sm_country_id" bson:"_sm_country_id"`
	Name        string             `json:"name" bson:"name"`
	Code        string             `json:"code" bson:"code"`
	Image       Limage             `json:"image" bson:"image"`
	Type        string             `json:"type" bson:"type"`
	CountriesID primitive.ObjectID `json:"_countries_id" bson:"_countries_id"`
	SeasonsID   primitive.ObjectID `json:"_seasons_id" bson:"_seasons_id"`
}

type Limage struct {
	Host     string `json:"host" bson:"host"`
	Path     string `json:"path" bson:"path"`
	Filename string `json:"filename" bson:"filename"`
}
