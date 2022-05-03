package handlers

import (
	"strconv"
	"work-request/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func (h *workRequestHandler) GetWorkRequestsForUser(c *fiber.Ctx) error {

	userTokenData, err := auth.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	page, err := strconv.Atoi(c.Params("page"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	number, err := strconv.Atoi(c.Params("number"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	status := c.Params("status")
	workRequests, err := h.workRequestApplication.GetWorkRequests(userTokenData.ID, status, page, number, "worker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(workRequests)
}
