package handlers

import (
	"projects/pkg/core/models"
	"projects/pkg/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *projectHandler) AddUserToProject(c *fiber.Ctx) error {

	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	addUserData := new(models.AddUserToProject)
	if err := c.BodyParser(&addUserData); err != nil {
		return fiber.ErrBadRequest
	}
	//TODO: check what happen if someone put a non existing type of the database
	document := c.Params("type")
	addUserData.OwnerId = userTokenData.ID
	if err := h.projectService.AddUserToProject(*addUserData, document); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)

}
