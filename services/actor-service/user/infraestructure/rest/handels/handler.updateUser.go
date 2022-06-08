package handlers

import (
	"actor-service/internal/auth"
	"actor-service/user/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	updatedUser := new(models.User)
	//TODO Change the parameters of this handler

	if err := c.BodyParser(&updatedUser); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := h.userApplication.UpdateUser(*updatedUser, userTokenData.ID); err != nil {
		return fiber.ErrBadRequest
	}
	return c.SendStatus(fiber.StatusOK)

}
