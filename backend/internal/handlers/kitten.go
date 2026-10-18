package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kitten-backend/internal/models"
	"kitten-backend/internal/repository"
	"strconv"
)

func GetKittens(c *fiber.Ctx) error {
	kittens, err := repository.GetAllKittens()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not get all kittens"})
	}

	return c.JSON(kittens)
}

func GetKitten(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	kitten, err := repository.GetKittenByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Kitten not found"})
	}

	return c.JSON(kitten)
}

func CreateKitten(c *fiber.Ctx) error {
	var kitten models.Kitten
	if err := c.BodyParser(&kitten); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Could not create kitten"})
	}

	if err := repository.CreateKitten(&kitten); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create kitten"})
	}

	return c.Status(fiber.StatusCreated).JSON(kitten)
}
