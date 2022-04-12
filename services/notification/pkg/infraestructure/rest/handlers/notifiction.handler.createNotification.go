package handlers

import (
	"notifications/pkg/core/models"
	"notifications/pkg/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *notificationHandler) CreateNotification(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	newNotificationData := new(models.Notification)
	if err := c.BodyParser(&newNotificationData); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	newNotificationData.Type = c.Params("type")
	newNotificationData.NotifierUserId = userTokenData.ID

	if err := h.notificationService.CreateNotification(*newNotificationData); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
