package database

import (
	"fmt"
	"log"
	"os"

	"kitten-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	log.Printf("Connecting to database with DSN: host=%s port=%s dbname=%s", host, port, dbname)

	var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successfully established")

	err = db.AutoMigrate(&models.Kitten{}, &models.User{})
	if err != nil {
		log.Printf("Database migration failed: %v", err)
	}
	log.Println("Database migrated successfully.")
	return db
}
