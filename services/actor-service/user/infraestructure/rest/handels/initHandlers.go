package handlers

import (
	"actor-service/user/application"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	UpdateUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
}
type userHandler struct {
	userApplication application.UserApplication
}

func NewUserHandler(userApplication application.UserApplication) UserHandler {
	return &userHandler{
		userApplication,
	}
}
