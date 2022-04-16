package handlers

import (
	"user/internal/auth"
	"user/pkg/core/models"

	"github.com/gofiber/fiber/v2"
)

//No puede ser vivible al front esta funcion solo deveria ser accesible desde el api getway
func (h *userHandler) AddProjectToUser(c *fiber.Ctx) error {

	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	addProjectData := new(models.AddProjectToUser)
	if err := c.BodyParser(&addProjectData); err != nil {
		return fiber.ErrBadRequest
	}
	addProjectData.User = userTokenData.ID
	//TODO: Check if what happen if someone put a tipe that doesn't exist on the db
	document := c.Params("type")

	if err := h.userService.AddProjectToUser(*addProjectData, document); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
