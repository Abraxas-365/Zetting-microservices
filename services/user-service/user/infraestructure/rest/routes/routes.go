package routes

import (
	"user/internal/auth"
	"user/user/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", handler.LoginUser)
	/*Register user*/
	users.Post("/register", handler.CreateUser)
	/*Check email exist*/
	// users.Get("/email=:email", handler.CheckEmailExist)
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
	/*get user by id*/
	users.Get("/id=:id", handler.GetUserById)
	/*Is user exist*/
	users.Get("/email=:email", handler.IsEmailInDb)
}
