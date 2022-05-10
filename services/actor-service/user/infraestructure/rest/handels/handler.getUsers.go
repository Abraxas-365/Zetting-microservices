package handlers

import "github.com/gofiber/fiber/v2"

func (h *userHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.userApplication.GetUsers()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)

}
