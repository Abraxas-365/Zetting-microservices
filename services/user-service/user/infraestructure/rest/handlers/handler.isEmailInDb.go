package handlers

import (
	"user/user/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) IsEmailInDb(c *fiber.Ctx) error {
	email := c.Params("email")
	if err := h.userApplication.IsUserExsist(models.Email(email)); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
