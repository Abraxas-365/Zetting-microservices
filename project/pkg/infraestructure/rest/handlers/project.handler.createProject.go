package handlers

import (
	"projects/pkg/core/models"
	"projects/pkg/internal/auth"

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
	newProjectId, err := h.projectService.CreateProject(createProjectData, userTokenData.ID)
	if err != nil {
		return c.SendStatus(fiber.ErrConflict.Code)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"msg": "correcto", "pid": newProjectId})

}
