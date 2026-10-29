package handlers

import (
	"kitten-backend/internal/auth"
	"kitten-backend/internal/response"
	"kitten-backend/internal/service"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	userService  *service.UserService
	emailService *service.EmailService
}

func NewAuthHandler(userService *service.UserService, emailService *service.EmailService) *AuthHandler {
	return &AuthHandler{
		userService:  userService,
		emailService: emailService,
	}
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

	if user.Role != "admin" {
		err := h.emailService.SendCode(c.Context(), user.Email)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"require_code": true,
			"message":      "Code sent to email",
		})
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

func (h *AuthHandler) LoginWithCode(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Code     string `json:"code"`
	}

	if err := c.BodyParser(&input); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid JSON", nil)
	}

	user, err := h.userService.GetUserByUsername(input.Username)
	if err != nil {
		user, err = h.userService.GetUserByEmail(input.Username)
		if err != nil {
			return response.SendError(c, fiber.StatusNotFound, "user not found", nil)
		}
	}

	err = h.emailService.VerifyCode(c.Context(), user.Email, input.Code)
	if err != nil {
		return response.SendError(c, fiber.StatusUnauthorized, "invalid or expired code", nil)
	}

	return h.generateTokens(c, user.ID, user.Username, user.Role)
}

func (h *AuthHandler) generateTokens(c *fiber.Ctx, userID uint, username, role string) error {
	accessToken, err := auth.GenerateToken(userID, username, role)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	refreshToken, err := auth.GenerateRefreshToken(userID, username, role)
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

	return c.JSON(fiber.Map{"access_token": accessToken, "role": role})
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&input); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid JSON", nil)
	}

	err := h.userService.Register(input.Username, input.Email, input.Password)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "user created"})
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

func (h *AuthHandler) MakeAdmin(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "invalid id", nil)
	}

	if err := h.userService.MakeAdmin(uint(id)); err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "internal error", nil)
	}

	return c.JSON(fiber.Map{"message": "User is now an admin"})
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		Path:     "/api/v1/refresh",
	})

	return c.JSON(fiber.Map{"message": "user logged out"})
}
