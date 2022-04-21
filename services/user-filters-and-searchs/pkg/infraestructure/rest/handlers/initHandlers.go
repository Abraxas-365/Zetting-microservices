package handlers

import (
	"user/pkg/application"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUsersByProfession(c *fiber.Ctx) error
	GetUsersByProject(c *fiber.Ctx) error
}
type userHandler struct {
	userApplication application.UserApplication
}

func NewUserHandler(UserApplication application.UserApplication) UserHandler {
	return &userHandler{
		UserApplication,
	}
}
