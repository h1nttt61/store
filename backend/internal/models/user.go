package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null" validate:"required,min=3"`
	Email    string `json:"email" gorm:"unique;not null" validate:"required,email"`
	Password string `json:"password" gorm:"not null" validate:"required,min=8"`
	Role     string `json:"role" gorm:"default:'user'"`
}
