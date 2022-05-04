package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *userHandler) GetUserById(c *fiber.Ctx) error {

	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	user, err := h.userApplication.GetUserById(userId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user.ExposeToPublic())
}
