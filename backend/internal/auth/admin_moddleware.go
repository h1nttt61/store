package auth

import (
	"github.com/gofiber/fiber/v2"
	"kitten-backend/internal/response"
)

func CheckAdmin(c *fiber.Ctx) error {
	role, ok := c.Locals("role").(string)
	if !ok || role != "admin" {
		return response.SendError(c, fiber.StatusForbidden, "Access denied: admins only", nil)
	}

	return c.Next()
}
