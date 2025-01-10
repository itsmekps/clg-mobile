package service

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	// "log"
)

// var UserServiceInstance *UserService

type UserService struct {
	repo *mongodb.UserRepository
}

func NewUserService(repo *mongodb.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// User-related methods
func (s *UserService) GetUser(id primitive.ObjectID) (*models.User, *errors.AppError) {

	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, errors.USER_NOT_FOUND
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*models.User, *errors.AppError) {

	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.USER_NOT_FOUND
	}
	return user, nil
}

// func (s *UserService) CreateUser(user *User) (*User, error) {
// 	// Implement logic using s.repo
// }

// func (s *UserService) UpdateUser(user *User) (*User, error) {
// 	// Implement logic using s.repo
// }

// func (s *UserService) DeleteUser(username string) error {
// 	// Implement logic using s.repo
// }
