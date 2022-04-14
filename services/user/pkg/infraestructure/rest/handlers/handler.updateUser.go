package handlers

import (
	"fmt"
	"user/internal/auth"
	"user/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	fmt.Println("---UpdateUser Route---")
	updatedUser := new(models.User)
	if err := c.BodyParser(&updatedUser); err != nil {
		return fiber.ErrBadRequest
	}
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	if err := h.userService.UpdateUser(updatedUser, userTokenData.ID); err != nil {
		return fiber.ErrBadRequest
	}
	return c.SendStatus(fiber.StatusOK)

}
