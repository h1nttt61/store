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

func (r *KittenRepository) GetAllKittens(limit, offset int, breed, sortBy string) ([]models.Kitten, error) {
	var kittens []models.Kitten
	query := r.db.Where(&models.Kitten{IsAvailable: true})

	if breed != "" {
		query = query.Where("breed = ?", breed)
	}
	if sortBy == "price_asc" {
		query = query.Order("price_asc")
	} else if sortBy == "price_desc" {
		query = query.Order("price_desc")
	} else {
		query = query.Order("created_at desc")
	}

	res := query.Limit(limit).Offset(offset).Find(&kittens)
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
