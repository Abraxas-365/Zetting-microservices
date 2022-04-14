package handlers

import "github.com/gofiber/fiber/v2"

func (h *projectHandler) GetProjectByProjectId(c *fiber.Ctx) error {

	projectId := c.Params("id")
	project, err := h.projectService.GetProjectByProjectId(projectId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(project)
}
