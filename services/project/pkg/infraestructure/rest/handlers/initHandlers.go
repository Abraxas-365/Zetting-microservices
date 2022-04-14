package handlers

import (
	"projects/pkg/core/ports"

	"github.com/gofiber/fiber/v2"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetMyProjects(c *fiber.Ctx) error
	GetProjectsWorkingOn(c *fiber.Ctx) error
	GetProjectByProjectId(c *fiber.Ctx) error
	AddUserToProject(c *fiber.Ctx) error
}
type projectHandler struct {
	projectService ports.ProjectService
}

func NewProjectHandler(projectService ports.ProjectService) ProjectHandler {
	return &projectHandler{
		projectService,
	}
}
