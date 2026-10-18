package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"kitten-backend/internal/database"
	"kitten-backend/internal/handlers"
)

func main() {
	_ = godotenv.Load()

	database.ConnectDB()
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api/v1")
	api.Get("/kittens", handlers.GetKittens)
	api.Get("/kittens/:id", handlers.GetKitten)
	api.Post("/kittens", handlers.CreateKitten)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
