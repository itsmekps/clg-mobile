// service/mongo/user_service.go
package service

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"
	"math"
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
func (s *PlayerService) GetPlayersList(page, limit int) ([]models.PlayerList, *Pagination, *errors.AppError) {
	// Fetch players from the repository
	players, totalCount, err := s.repo.PlayerList(page, limit)
	if err != nil {
		return nil, nil, errors.INTERNAL_SERVER_ERROR
	}

	// Calculate pagination metadata
	totalPages := int(math.Ceil(float64(totalCount) / float64(limit)))
	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = 0 // No next page
	}
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 0 // No previous page
	}

	// Create pagination object
	pagination := &Pagination{
		CurrentPage: page,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		TotalPages:  totalPages,
	}

	return players, pagination, nil
}

func (s *PlayerService) SearchPlayers(query string, limit int) ([]models.PlayerSearch, *errors.AppError) {
	// Fetch players matching the search query from the repository
	players, err := s.repo.SearchPlayers(query, limit)
	if err != nil {
		return nil, errors.INTERNAL_SERVER_ERROR
	}

	return players, nil
}

// Pagination struct to hold pagination metadata
type Pagination struct {
	CurrentPage int `json:"current_page"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
	TotalPages  int `json:"total_pages"`
}
