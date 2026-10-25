package service

import (
	"kitten-backend/internal/models"
)

type KittenRepository interface {
	GetAllKittens(limit, offset int, breed, sortBy string) ([]models.Kitten, error)
	GetKittenByID(id uint) (models.Kitten, error)
	CreateKitten(kitten *models.Kitten) error
}

type KittenService struct {
	repo KittenRepository
}

func NewKittenService(repo KittenRepository) *KittenService {
	return &KittenService{repo: repo}
}

func (s *KittenService) GetAllKittens(limit, offset int, breed, sortBy string) ([]models.Kitten, error) {
	return s.repo.GetAllKittens(limit, offset, breed, sortBy)
}

func (s *KittenService) GetKittenByID(id uint) (models.Kitten, error) {
	return s.repo.GetKittenByID(id)
}

func (s *KittenService) CreateKitten(kitten *models.Kitten) error {
	return s.repo.CreateKitten(kitten)
}
