package repository

import (
	"kitten-backend/internal/models"

	"gorm.io/gorm"
)

type KittenRepository struct {
	db *gorm.DB
}

func NewKittenRepository(db *gorm.DB) *KittenRepository {
	return &KittenRepository{db: db}
}

func (r *KittenRepository) GetAllKittens(limit int, offset int) ([]models.Kitten, error) {
	var kittens []models.Kitten
	res := r.db.Where("is_available = ?", true).Find(&kittens).Limit(limit).Offset(offset).Find(&kittens)
	return kittens, res.Error
}

func (r *KittenRepository) GetKittenByID(id uint) (models.Kitten, error) {
	var kitten models.Kitten
	res := r.db.First(&kitten, id)
	return kitten, res.Error
}

func (r *KittenRepository) CreateKitten(kitten *models.Kitten) error {
	res := r.db.Create(&kitten)
	return res.Error
}
