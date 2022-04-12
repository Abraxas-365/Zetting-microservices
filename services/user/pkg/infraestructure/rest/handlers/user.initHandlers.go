package handlers

import (
	"user/pkg/core/ports"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
	GetUserById(c *fiber.Ctx) error
	CheckEmailExist(c *fiber.Ctx) error
	GetUsersByProfession(c *fiber.Ctx) error
	GetUsersByProject(c *fiber.Ctx) error
	AddProjectToUser(c *fiber.Ctx) error
	UploadProfileImage(c *fiber.Ctx) error
}
type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) UserHandler {
	return &userHandler{
		userService,
	}
}
