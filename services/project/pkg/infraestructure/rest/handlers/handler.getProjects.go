package handlers

import (
	"projects/internal/auth"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *projectHandler) GetMyProjects(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	page, _ := strconv.Atoi(c.Params("page"))
	myProjects, err := h.projectService.GetProjects(userTokenData.ID, "owners", page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(myProjects)

}

func (h *projectHandler) GetProjectsWorkingOn(c *fiber.Ctx) error {
	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	page, _ := strconv.Atoi(c.Params("page"))
	projects, err := h.projectService.GetProjects(userTokenData.ID, "workers", page)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(projects)

}
