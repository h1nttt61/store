package handlers

import (
	"github.com/gofiber/fiber/v2"
	"kitten-backend/internal/auth"
	"kitten-backend/internal/response"
	"kitten-backend/internal/service"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid JSON", nil)
	}

	user, err := h.userService.Login(input.Username, input.Password)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "invalid username or password", nil)
	}

	token, err := auth.GenerateToken(user.ID, user.Username)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	return c.JSON(fiber.Map{"token": token})
}
