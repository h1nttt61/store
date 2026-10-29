package auth

import (
	"kitten-backend/internal/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CheckJWT(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return response.SendError(c, fiber.StatusUnauthorized, "Missing or invalid token", nil)
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return response.SendError(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}
	tokenString := parts[1]
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return response.SendError(c, fiber.StatusUnauthorized, "Invalid token", nil)
	}

	c.Locals("userID", claims.UserID)
	c.Locals("username", claims.Username)
	c.Locals("role", claims.Role)

	return c.Next()
}
