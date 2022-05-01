package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *projectHandler) GetProjectByProjectId(c *fiber.Ctx) error {

	projectId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	project, err := h.projectApplication.GetProjectByProjectId(projectId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(project)
}
