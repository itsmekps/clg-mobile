package service

import (
	"fiber-boilerplate/app/models"
	"fiber-boilerplate/app/repository/mongodb"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewsService struct {
	Repo *mongodb.NewsRepository
}

func NewNewsService(repo *mongodb.NewsRepository) *NewsService {
	return &NewsService{Repo: repo}
}

func (s *NewsService) CreateNews(news *models.News) error {
	return s.Repo.CreateNews(news)
}

func (s *NewsService) GetNewsByID(id primitive.ObjectID) (*models.News, error) {
	return s.Repo.GetNewsByID(id)
}

func (s *NewsService) UpdateNews(id primitive.ObjectID, update map[string]interface{}) error {
	return s.Repo.UpdateNews(id, update)
}

func (s *NewsService) DeleteNews(id primitive.ObjectID) error {
	return s.Repo.DeleteNews(id)
}

func (s *NewsService) ListNews() ([]models.News, error) {
	return s.Repo.ListNews()
}
