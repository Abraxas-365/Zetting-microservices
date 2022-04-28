package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *workRequestHandler) GetWorkRequestsByWorker(c *fiber.Ctx) error {

	// userTokenData, err := auth.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }
	page, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	number, err := strconv.Atoi(c.Params("number"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	status := c.Params("status")
	workerId, err := uuid.Parse(c.Params("worker_id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	workRequests, err := h.workRequestApplication.GetWorkRequests(workerId, status, page, number, "worker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(workRequests)
}
