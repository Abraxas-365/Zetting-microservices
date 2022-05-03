package handlers

import (
	"projects/project/application"

	"github.com/gofiber/fiber/v2"
)

type ProjectHandler interface {
	CreateProject(c *fiber.Ctx) error
	GetMyProjects(c *fiber.Ctx) error
	GetProjectsWorkingOn(c *fiber.Ctx) error
	GetProjectByProjectId(c *fiber.Ctx) error
}
type projectHandler struct {
	projectApplication application.ProjectApplication
}

func NewProjectHandler(projectApplication application.ProjectApplication) ProjectHandler {
	return &projectHandler{
		projectApplication,
	}
}
