package handlers

import (
	"notifications/notification/application"

	"github.com/gofiber/fiber/v2"
)

type NotificationsHandler interface {
	GetNotifications(c *fiber.Ctx) error
	GetCompleteNotification(c *fiber.Ctx) error
}
type notificationHandler struct {
	notificationApplication application.NotificationApplication
}

func NewNotificationHandler(notificationApplication application.NotificationApplication) NotificationsHandler {
	return &notificationHandler{
		notificationApplication,
	}
}
