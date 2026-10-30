package main

import (
	"context"
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
	"github.com/redis/go-redis/v9"
)

func main() {
	_ = godotenv.Load()

	validator.Init()

	db := database.ConnectDB()
	rdb := redis.NewClient(&redis.Options{Addr: os.Getenv("REDIS_HOST") + ":6379"})
	kittenRepo := repository.NewKittenRepository(db)
	kittenService := service.NewKittenService(kittenRepo)
	kittenHandler := handlers.NewKittenHandler(kittenService)
	uploadHander := handlers.NewUploadHandler()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	emailService := service.NewEmailService(rdb)
	userHandler := handlers.NewAuthHandler(userService, emailService)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatal("Failed to connect to Redis: %v", err)
	}

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	app.Static("/uploads", "./uploads")

	api := app.Group("/api/v1")
	api.Get("/kittens", kittenHandler.GetKittens)
	api.Get("/kittens/:id", kittenHandler.GetKitten)
	api.Post("/login", userHandler.Login)
	api.Post("/register", userHandler.Register)
	api.Post("/refresh", userHandler.Refresh)
	api.Post("/login-with-code", userHandler.LoginWithCode)
	api.Post("/logout", userHandler.Logout)

	protectedAdmin := api.Group("/", auth.CheckJWT, auth.CheckAdmin)

	protectedAdmin.Post("/kittens", kittenHandler.CreateKitten)
	protectedAdmin.Post("/upload", uploadHander.UploadImage)
	protectedAdmin.Post("/users/:id/make-admin", userHandler.MakeAdmin)
	protectedAdmin.Get("/users", userHandler.GetUsers)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
