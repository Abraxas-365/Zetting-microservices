package routes

import (
	"projects/user/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App, handler handlers.UserHandler) {

	project := app.Group("/api/projects/users")

	/*get user by project id*/
	project.Get("/id=:id", handler.GetUserById)
}
