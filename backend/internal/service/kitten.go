package service

import (
	"kitten-backend/internal/models"
	"kitten-backend/internal/repository"
)

type KittenService struct {
	repo *repository.KittenRepository
}

func NewKittenService(repo *repository.KittenRepository) *KittenService {
	return &KittenService{repo: repo}
}

func (s *KittenService) GetAllKittens(limit, offset int) ([]models.Kitten, error) {
	return s.repo.GetAllKittens(limit, offset)
}

func (s *KittenService) GetKittenByID(id uint) (models.Kitten, error) {
	return s.repo.GetKittenByID(id)
}

func (s *KittenService) CreateKitten(kitten *models.Kitten) error {
	return s.repo.CreateKitten(kitten)
}
