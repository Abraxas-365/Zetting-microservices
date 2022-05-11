package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) GetUsers(c *fiber.Ctx) error {

	page, err := strconv.Atoi(c.Params("page", "1"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	users, err := h.userApplication.GetUsers(page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)

}
