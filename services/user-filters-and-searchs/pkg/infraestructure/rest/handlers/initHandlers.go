package handlers

import (
	"user/pkg/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsersByProfession(c *fiber.Ctx) error
	GetUsersByProject(c *fiber.Ctx) error
}
type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) UserHandler {
	return &userHandler{
		userService,
	}
}
