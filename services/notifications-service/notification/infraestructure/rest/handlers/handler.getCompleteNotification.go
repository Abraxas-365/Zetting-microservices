package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *notificationHandler) GetCompleteNotification(c *fiber.Ctx) error {

	notificationId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	notification, err := h.notificationApplication.GetCompleteNotification(notificationId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(notification)
}
