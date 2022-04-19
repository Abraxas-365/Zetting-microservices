package handlers

import (
	"work-request/pkg/core/ports"

	"github.com/gofiber/fiber/v2"
)

type WorkRequestHandler interface {
	CreateWorkRequest(c *fiber.Ctx) error
	GetWorkRequestsByWorker(c *fiber.Ctx) error
	GetWorkRequestsByProject(c *fiber.Ctx) error
}
type workRequestHandler struct {
	workRequestService ports.WorkRequestService
}

func NewWorkRequestHandler(workRequestService ports.WorkRequestService) WorkRequestHandler {
	return &workRequestHandler{
		workRequestService,
	}
}
