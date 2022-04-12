package handlers

import (
	"work-request/pkg/core/models"
	"work-request/pkg/internal/auth"

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
	newWorkrequestData.OwnerId = userTokenData.ID
	newWorkrequestData, err = h.workRequestService.CreateWorkRequest(*newWorkrequestData)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(newWorkrequestData)

}
