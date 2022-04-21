package handlers

import (
	"projects/internal/auth"
	"projects/pkg/core/models"

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
	addUserData.Owner = userTokenData.ID
	if err := h.projectApplication.AddUserToProject(*addUserData, document); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)

}
