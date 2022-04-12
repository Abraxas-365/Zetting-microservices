package handlers

import (
	"fmt"
	"user/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

// Register.
// @Register
// @Summary  register.
// @Tags     Auth
// @Accept   json
// @Produce  json
// @Param    register  body      models.UserRegistration  true  "Register"
// @Success  200    {object}  models.AuthUser
// @Router   /users/register [post]

func (h *userHandler) CreateUser(c *fiber.Ctx) error {
	fmt.Println("---Register Route---")
	userRegisterData := new(models.User)
	if err := c.BodyParser(&userRegisterData); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	fmt.Println(userRegisterData)

	Newuser, err := h.userService.CreateUser(*userRegisterData)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	authUser, err := h.userService.LoginUser(Newuser.Contact.Email, Newuser.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(authUser)

}
