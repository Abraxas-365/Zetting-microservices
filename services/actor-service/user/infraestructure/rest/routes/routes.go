package routes

import (
	"actor-service/internal/auth"
	"actor-service/user/infraestructure/rest/handels"

	"github.com/gofiber/fiber/v2"
)

func UsersRoute(app *fiber.App, handler handlers.UserHandler) {
	users := app.Group("/api/actor")
	users.Put("/update", auth.JWTProtected(), handler.UpdateUser)
}
