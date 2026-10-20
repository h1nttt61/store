package response

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	Error   string            `json:"error"`
	Details map[string]string `json:"details,omitempty"`
}

func SendError(c *fiber.Ctx, status int, message string, details map[string]string) error {
	return c.Status(status).JSON(ErrorResponse{
		Error:   message,
		Details: details,
	})
}
