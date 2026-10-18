package repository

import (
	"kitten-backend/internal/database"
	"kitten-backend/internal/models"
)

func GetAllKittens() ([]models.Kitten, error) {
	var Kittens []models.Kitten
	result := database.DB.Where("is_available = ?", true).Find(&Kittens)
	return Kittens, result.Error
}

func GetKittenByID(id uint) (models.Kitten, error) {
	var Kitten models.Kitten
	result := database.DB.First(&Kitten, id)
	return Kitten, result.Error
}

func CreateKitten(kitten *models.Kitten) error {
	result := database.DB.Create(&kitten)
	return result.Error
}
