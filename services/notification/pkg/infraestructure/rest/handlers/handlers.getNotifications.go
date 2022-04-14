package handlers

import (
	"notifications/internal/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *notificationHandler) GetNotifications(c *fiber.Ctx) error {

	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	page, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	myNotifications, err := h.notificationService.GetNotifications(userTokenData.ID, page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(myNotifications)
}
