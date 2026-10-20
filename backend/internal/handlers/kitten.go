package handlers

import (
	val "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"kitten-backend/internal/models"
	"kitten-backend/internal/repository"
	"kitten-backend/internal/response"
	"kitten-backend/internal/validator"
	"strconv"
	"strings"
)

type KittenHandler struct {
	repo *repository.KittenRepository
}

func NewKittenHandler(repo *repository.KittenRepository) *KittenHandler {
	return &KittenHandler{repo: repo}
}

func (h *KittenHandler) GetKittens(c *fiber.Ctx) error {
	kittens, err := h.repo.GetAllKittens()
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not get all kittens", nil)
	}

	return c.JSON(kittens)
}

func (h *KittenHandler) GetKitten(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Could not parse id", nil)
	}

	kitten, err := h.repo.GetKittenByID(uint(id))
	if err != nil {
		return response.SendError(c, fiber.StatusNotFound, "Kitten not found", nil)
	}

	return c.JSON(kitten)
}

func (h *KittenHandler) CreateKitten(c *fiber.Ctx) error {
	var kitten models.Kitten
	if err := c.BodyParser(&kitten); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid JSON format", nil)
	}

	if err := validator.Validate.Struct(kitten); err != nil {
		errorDetails := make(map[string]string)

		for _, err := range err.(val.ValidationErrors) {
			field := strings.ToLower(err.Field())
			errorDetails[field] = "Invalid value for " + err.Field()

			if err.Tag() == "gt" {
				errorDetails[field] = "Must be greater than " + err.Param()
			}
			if err.Tag() == "required" {
				errorDetails[field] = "This field is required"
			}
		}

		return response.SendError(c, fiber.StatusBadRequest, "Validation failed", errorDetails)
	}

	if err := h.repo.CreateKitten(&kitten); err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not create kitten", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(kitten)
}
