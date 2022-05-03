package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *workRequestHandler) GetWorkRequestsByProject(c *fiber.Ctx) error {

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
	projectId, err := uuid.Parse(c.Params("project_id"))
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	workRequests, err := h.workRequestApplication.GetWorkRequests(projectId, status, page, number, "project._id")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(workRequests)

}
