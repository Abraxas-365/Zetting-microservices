package handlers

import (
	"projects/user/application"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUserById(c *fiber.Ctx) error
}
type userHandler struct {
	userApp application.UserApplication
}

func NewUserHandler(userApp application.UserApplication) UserHandler {
	return &userHandler{
		userApp,
	}
}
