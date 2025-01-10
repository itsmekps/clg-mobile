package service

import (
	"fiber-boilerplate/app/errors"
	"fiber-boilerplate/app/repository/mongodb"
	"fiber-boilerplate/utils"
)

type AuthService struct {
	repo *mongodb.UserRepository
}

func NewAuthService(repo *mongodb.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Login(email, password string) (string, *errors.AppError) {
	// Fetch user from repository
	user, err := s.repo.GetUserByEmail(email)
	if err != nil || user == nil {
		return "", errors.USER_NOT_FOUND
	}

	// Validate password
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.INVALID_PASSWORD
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", errors.INTERNAL_SERVER_ERROR
	}

	return token, nil
}
