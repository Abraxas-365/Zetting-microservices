package handlers

import (
	"notifications/pkg/core/ports"

	"github.com/gofiber/fiber/v2"
)

type NotificationsHandler interface {
	GetNotifications(c *fiber.Ctx) error
	GetCompleteNotification(c *fiber.Ctx) error
}
type notificationHandler struct {
	notificationService ports.NotificationService
}

func NewNotificationHandler(notificationService ports.NotificationService) NotificationsHandler {
	return &notificationHandler{
		notificationService,
	}
}
