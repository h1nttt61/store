package main

import (
	"kitten-backend/internal/auth"
	"kitten-backend/internal/repository"
	"kitten-backend/internal/service"
	"log"
	"os"

	"kitten-backend/internal/database"
	"kitten-backend/internal/handlers"

	"kitten-backend/internal/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	validator.Init()

	db := database.ConnectDB()
	kittenRepo := repository.NewKittenRepository(db)
	kittenService := service.NewKittenService(kittenRepo)
	kittenHandler := handlers.NewKittenHandler(kittenService)
	uploadHander := handlers.NewUploadHandler()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewAuthHandler(userService)

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Static("/uploads", "./uploads")

	api := app.Group("/api/v1")
	api.Get("/kittens", kittenHandler.GetKittens)
	api.Get("/kittens/:id", kittenHandler.GetKitten)
	api.Post("/login", userHandler.Login)

	protected := app.Group("/", auth.CheckJWT)

	protected.Post("/kittens", kittenHandler.CreateKitten)
	protected.Post("/upload", uploadHander.UploadImage)

	api.Get("/")

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
