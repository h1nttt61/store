package handlers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"kitten-backend/internal/response"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

func (h *UploadHandler) UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "No file upload", nil)
	}

	if file.Size > 5*1024*1024 {
		return response.SendError(c, fiber.StatusBadRequest, "File too large(max 5 mb)", nil)
	}

	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	savepath := filepath.Join("./uploads", newFilename)

	if err := c.SaveFile(file, savepath); err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Could not save file", nil)
	}

	return c.JSON(fiber.Map{"url": "/uploads/" + newFilename})
}
