package routes

import (
	"user/internal/auth"
	"user/pkg/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	users := app.Group("/api/users")
	/*Update user*/
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
	/*Add Project id to user*/
	users.Put("/add_project", auth.JWTProtected(), handler.AddProjectToUser)
}
