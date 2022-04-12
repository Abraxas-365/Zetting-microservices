package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func (h *workRequestHandler) GetWorkRequestsById(c *fiber.Ctx) error {

	// userTokenData, err := auth.ExtractTokenMetadata(c)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }
	workRequestId := c.Params("id")
	workRequest, err := h.workRequestService.GetWorkRequest(workRequestId)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(workRequest)
}
