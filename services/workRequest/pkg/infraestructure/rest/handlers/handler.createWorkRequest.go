package handlers

import (
	"work-request/internal/auth"
	"work-request/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *workRequestHandler) CreateWorkRequest(c *fiber.Ctx) error {
	newWorkrequestData := new(models.WorkRequest)
	if err := c.BodyParser(&newWorkrequestData); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	newWorkrequestData.Owner.ID = userTokenData.ID
	if err := h.workRequestApplication.CreateWorkRequest(*newWorkrequestData); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)

}
