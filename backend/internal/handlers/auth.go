package handlers

import (
	"kitten-backend/internal/auth"
	"kitten-backend/internal/response"
	"kitten-backend/internal/service"
	"time"

	"github.com/gofiber/fiber/v2"
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

	accessToken, err := auth.GenerateToken(user.ID, user.Username, user.Role)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	refreshToken, err := auth.GenerateRefreshToken(user.ID, user.Username, user.Role)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/api/v1/refresh",
	})

	return c.JSON(fiber.Map{"access_token": accessToken, "role": user.Role})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid JSON", nil)
	}

	err := h.userService.Register(input.Username, input.Password)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"meesage": "admin created"})
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return response.SendError(c, fiber.StatusBadRequest, "invalid refresh token", nil)
	}

	claims, err := auth.ValidateToken(refreshToken)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid refresh token", nil)
	}

	newAccessToken, err := auth.GenerateToken(claims.UserID, claims.Username, claims.Role)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	return c.JSON(fiber.Map{"access_token": newAccessToken})
}
