package handlers

import (
	"work-request/workrequest/application"

	"github.com/gofiber/fiber/v2"
)

type WorkRequestHandler interface {
	CreateWorkRequest(c *fiber.Ctx) error
	GetWorkRequestsForUser(c *fiber.Ctx) error
	GetWorkRequestsByProject(c *fiber.Ctx) error
	AnswerWorkRequest(c *fiber.Ctx) error
}
type workRequestHandler struct {
	workRequestApplication application.WorkRequestApplication
}

func NewWorkRequestHandler(workRequestApplication application.WorkRequestApplication) WorkRequestHandler {
	return &workRequestHandler{
		workRequestApplication,
	}
}
