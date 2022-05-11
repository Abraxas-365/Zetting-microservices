package handlers

import (
	"actor-service/user/core/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) FilterUsers(c *fiber.Ctx) error {
	filter := new(models.Filter)
	if err := c.BodyParser(&filter); err != nil {
		return fiber.ErrBadRequest
	}
	page, err := strconv.Atoi(c.Params("page", "1"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	users, err := h.userApplication.FilterUsers(*filter, page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)

}
