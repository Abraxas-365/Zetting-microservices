package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
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
	workerId := c.Params("worker_id")
	workRequests, err := h.workRequestService.GetWorkRequests(workerId, status, page, number, "worker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(workRequests)
}
