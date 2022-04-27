package handlers

import "github.com/gofiber/fiber/v2"

func (h *userHandler) GetUserById(c *fiber.Ctx) error {

	userId := c.Params("id")
	user, err := h.userApplication.GetUserById(userId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
