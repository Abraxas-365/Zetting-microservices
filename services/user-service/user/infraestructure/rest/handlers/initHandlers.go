package handlers

import (
	"user/user/application"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	UpdateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	CreateUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
}
type userHandler struct {
	userApplication application.UserApplication
}

func NewUserHandler(userApplication application.UserApplication) UserHandler {
	return &userHandler{
		userApplication,
	}
}
