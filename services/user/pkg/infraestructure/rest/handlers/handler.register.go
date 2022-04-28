package handlers

import (
	"user/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	newUser := new(models.User)
	if err := c.BodyParser(&newUser); err != nil {
		return fiber.ErrBadRequest
	}
	user, token, err := h.userApplication.CreateUser(newUser.New())
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user, "token": token})

}
