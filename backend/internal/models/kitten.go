package models

import "gorm.io/gorm"

type Kitten struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null" validate:"required,min=2,max=50"`
	Breed       string  `json:"breed" gorm:"not null" validate:"required"`
	AgeMonths   int     `json:"age_months" validate:"gte=0"`
	Price       float64 `json:"price" gorm:"not null" validate:"required,gt=0"`
	Description string  `json:"description" validate:"omitempty,max=500"`
	ImageURL    string  `json:"image_url" validate:"omitempty,url"`
	IsAvailable bool    `json:"is_available" gorm:"default:true"`
}
