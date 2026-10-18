package models

import (
	"gorm.io/gorm"
)

type Kitten struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Breed       string  `json:"breed" gorm:"not null"`
	AgeMonths   int     `json:"age_months"`
	Price       float64 `json:"price" gorm:"not null`
	Description string  `json:"description"`
	ImageURL    string  `json:"image_url"`
	IsAvailable bool    `json:"is_available" gorm:"default:true"`
}
