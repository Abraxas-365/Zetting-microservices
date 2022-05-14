package routes

import (
	"actor-service/internal/auth"
	handlers "actor-service/user/infraestructure/rest/handels"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	users := app.Group("/api/actor")
	/*update user*/
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
	/*get all users*/
	users.Get("/page=:page", handler.GetUsers)
	/*get user with filters*/
	users.Get("/filter/page=:page", handler.FilterUsers)
	/*get user by id*/
	users.Get("/id=:id", handler.GetUserById)
}
