package main

import (
	"kitten-backend/internal/repository"
	"log"
	"os"

	"kitten-backend/internal/database"
	"kitten-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"kitten-backend/internal/validator"
)

func main() {
	_ = godotenv.Load()

	validator.Init()

	db := database.ConnectDB()
	kittenRepo := repository.NewKittenRepository(db)
	kittenHandler := handlers.NewKittenHandler(kittenRepo)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/kittens", kittenHandler.GetKittens)
	api.Get("/kittens/:id", kittenHandler.GetKitten)
	api.Post("/kittens", kittenHandler.CreateKitten)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
