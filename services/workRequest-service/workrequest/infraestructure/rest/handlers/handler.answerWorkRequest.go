package handlers

import (
	"work-request/internal/auth"
	"work-request/workrequest/core/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h workRequestHandler) AnswerWorkRequest(c *fiber.Ctx) error {

	body := struct {
		ID     string        `json:"id,omitemptyo"`
		Status models.Status `json:"status,omitempty"`
	}{}

	// new(models.WorkRequest)
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	id, err := uuid.Parse(body.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := h.workRequestApplication.AnswerWorkRequest(id, userTokenData.ID, body.Status); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
