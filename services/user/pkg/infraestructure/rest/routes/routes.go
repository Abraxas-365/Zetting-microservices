package routes

import (
	"user/internal/auth"
	"user/pkg/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	/*SERVE*/
	static := app.Group("/static")
	static.Static("/app_default_images", "./static/app_default_images", fiber.Static{
		Browse: false, // default: false
	})
	users := app.Group("/api/users")
	/*Login user*/
	users.Post("/login", handler.LoginUser)
	/*Register user*/
	users.Post("/register", handler.CreateUser)
	/*Check email exist*/
	users.Get("/email=:email", handler.CheckEmailExist)
	/*Get user by Project id*/
	users.Get("/project_id=:project_id", auth.JWTProtected(), handler.GetUsersByProject)
	/*Get user */
	users.Get("/id=:id", handler.GetUserById)
	/*Get users by profession*/
	users.Get("/profession=:profession/page=:page", auth.JWTProtected(), handler.GetUsersByProfession)
	/*Update user*/
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
	/*Upload file image*/
	users.Put("/upload/perfil_image", auth.JWTProtected(), handler.UploadProfileImage)
	/*Add Project id to user*/
	users.Put("/add_project", auth.JWTProtected(), handler.AddProjectToUser)
}
