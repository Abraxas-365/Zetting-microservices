package handlers

import (
	"work-request/internal/auth"
	"work-request/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h workRequestHandler) AnswerWorkRequest(c *fiber.Ctx) error {

	answerWorkrequestData := new(models.WorkRequest)
	if err := c.BodyParser(&answerWorkrequestData); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	answerWorkrequestData.Worker = userTokenData.ID
	if err := h.workRequestService.AnswerWorkRequest(*answerWorkrequestData); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
