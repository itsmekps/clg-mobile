// service/mongo/user_service.go
package service

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
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
func (s *PlayerService) GetPlayersList(page, limit int) ([]models.PlayerList, *errors.AppError) {

	playerList, err := s.repo.GetPlayersList(page, limit)
	if err != nil {
		return nil, errors.INTERNAL_SERVER_ERROR
	}
	return playerList, nil

	// return s.repo.GetPlayersList(page, limit)
}
