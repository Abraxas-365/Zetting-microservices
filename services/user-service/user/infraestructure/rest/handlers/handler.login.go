package handlers

import (
	"user/user/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	body := struct {
		Email    models.Email    `json:"email"`
		Password models.Password `json:"password"`
	}{}
	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}
	user, token, err := h.userApplication.LoginUser(body.Email, body.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"user": user.ExposeToPublic(), "token": token})
}
