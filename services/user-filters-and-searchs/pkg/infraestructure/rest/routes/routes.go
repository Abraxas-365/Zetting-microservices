package routes

import (
	"user/internal/auth"
	"user/pkg/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	users := app.Group("/api/users-search")
	/*Get user by Project id*/
	users.Get("/project_id=:project_id", auth.JWTProtected(), handler.GetUsersByProject)
	/*Get users by profession*/
	users.Get("/profession=:profession/page=:page", auth.JWTProtected(), handler.GetUsersByProfession)
}
