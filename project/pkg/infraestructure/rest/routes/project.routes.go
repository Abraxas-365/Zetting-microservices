package routes

import (
	"projects/pkg/infraestructure/rest/handlers"
	"projects/pkg/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func ProjectsRoute(app *fiber.App, handler handlers.ProjectHandler) {

	project := app.Group("/api/projects")

	/*Create a new project*/
	project.Post("/new", auth.JWTProtected(), handler.CreateProject)
	/*Get my projects*/
	project.Get("/myprojects/page=:page", auth.JWTProtected(), handler.GetMyProjects)
	/*get projects im in*/
	project.Get("/projects/page=:page", auth.JWTProtected(), handler.GetProjectsWorkingOn)
	/*ger project by project id*/
	project.Get("/id=:id", auth.JWTProtected(), handler.GetProjectByProjectId)
	/*AddUserToProject*/
	project.Post("/adduser/type:type", auth.JWTProtected(), handler.AddUserToProject)
}
