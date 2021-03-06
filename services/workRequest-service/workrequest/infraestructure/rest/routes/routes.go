package routes

import (
	"work-request/internal/auth"
	"work-request/workrequest/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func WorkRequestRoute(app *fiber.App, handler handlers.WorkRequestHandler) {

	workRequest := app.Group("/api/work-request")

	/*Create a new project*/
	workRequest.Post("/new", auth.JWTProtected(), handler.CreateWorkRequest)
	/*Get all the work request of a project*/
	workRequest.Get("/project_id=:project_id/status=:status/page=:page/number=:number", auth.JWTProtected(), handler.GetWorkRequestsByProject)
	/*Get all the work request a user have resive*/
	workRequest.Get("/status=:status/page=:page/number=:number", auth.JWTProtected(), handler.GetWorkRequestsForUser)
	/*answer work request*/
	workRequest.Post("/answer", auth.JWTProtected(), handler.AnswerWorkRequest)
}
