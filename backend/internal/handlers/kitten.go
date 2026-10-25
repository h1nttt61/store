package handlers

import (
	"kitten-backend/internal/models"
	"kitten-backend/internal/response"
	"kitten-backend/internal/service"
	"kitten-backend/internal/validator"
	"strconv"
	"strings"

	val "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type KittenHandler struct {
	service *service.KittenService
}

func NewKittenHandler(service *service.KittenService) *KittenHandler {
	return &KittenHandler{service: service}
}

func (h *KittenHandler) GetKittens(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	page := c.QueryInt("page", 1)
	if page <= 0 {
		page = 1
	}

	offset := (page - 1) * limit

	breed := c.Query("breed")
	sortBy := c.Query("sortBy")

	kittens, err := h.service.GetAllKittens(limit, offset, breed, sortBy)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not get all kittens", nil)
	}

	return c.JSON(fiber.Map{
		"data":  kittens,
		"page":  page,
		"limit": limit,
	})
}

func (h *KittenHandler) GetKitten(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Could not parse id", nil)
	}

	kitten, err := h.service.GetKittenByID(uint(id))
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

	if err := h.service.CreateKitten(&kitten); err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not create kitten", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(kitten)
}
