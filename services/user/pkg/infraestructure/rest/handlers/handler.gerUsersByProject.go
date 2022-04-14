package handlers

import "github.com/gofiber/fiber/v2"

func (h *userHandler) GetUsersByProject(c *fiber.Ctx) error {
	projectId := c.Params("project_id")

	//TODO: ADD ownwer to document
	users, err := h.userService.GetUsersByProject(projectId, "worker")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(users)
}
