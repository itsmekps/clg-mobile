// service/mongo/user_service.go
package service

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	// "log"
)

// var PlayerServiceInstance *PlayerService

type PlayerService struct {
	repo *mongodb.PlayerRepository
}

func NewPlayerService(repo *mongodb.PlayerRepository) *PlayerService {
	return &PlayerService{repo: repo}
}

// User-related methods
func (s *PlayerService) GetPlayersList(page, limit int) ([]models.Player, *errors.AppError) {

	fields := bson.M{
		"fullname":    1,
		"dateofbirth": 1,
		"_id":         0, // Exclude the _id field
	}
	playerList, err := s.repo.FindAll(fields, page, limit)
	if err != nil {
		return nil, errors.INTERNAL_SERVER_ERROR
	}
	return playerList, nil

	// return s.repo.GetPlayersList(page, limit)
}
