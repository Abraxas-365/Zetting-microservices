package handlers

import "github.com/gofiber/fiber/v2"

func (h *notificationHandler) GetCompleteNotification(c *fiber.Ctx) error {

	notificationId := c.Params("id")
	notification, err := h.notificationApplication.GetCompleteNotification(notificationId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(notification)
}
