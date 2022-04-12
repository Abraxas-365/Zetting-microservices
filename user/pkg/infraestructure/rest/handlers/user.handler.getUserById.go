package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetUser.
// @GetUser
// @Summary   get user.
// @Tags    User
// @Success  200
// @Security  ApiKeyAuth
// @Router    /users/ [get]

func (h *userHandler) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("id")
	fmt.Println(userId)
	user, err := h.userService.GetUserById(userId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(user)

}
