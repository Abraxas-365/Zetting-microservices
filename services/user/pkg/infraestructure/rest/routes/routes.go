package routes

import (
	"user/internal/auth"
	"user/pkg/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	/*SERVE*/
	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", handler.LoginUser)
	/*Register user*/
	users.Post("/register", handler.CreateUser)
	/*Check email exist*/
	// users.Get("/email=:email", handler.CheckEmailExist)
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
}
