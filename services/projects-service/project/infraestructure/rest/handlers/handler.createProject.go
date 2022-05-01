package handlers

import (
	"projects/internal/auth"
	"projects/project/core/models"

	"github.com/gofiber/fiber/v2"
)

func (h *projectHandler) CreateProject(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	createProjectData := new(models.Project)
	if err := c.BodyParser(&createProjectData); err != nil {
		return fiber.ErrBadRequest
	}
	newProject, err := h.projectApplication.CreateProject(createProjectData.New(userTokenData.ID))
	if err != nil {
		return c.SendStatus(fiber.ErrConflict.Code)
	}
	return c.Status(fiber.StatusOK).JSON(newProject)

}
