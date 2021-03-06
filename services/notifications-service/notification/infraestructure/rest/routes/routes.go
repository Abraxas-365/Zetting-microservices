package routes

import (
	"notifications/internal/auth"
	"notifications/notification/infraestructure/rest/handlers"

	"github.com/gofiber/fiber/v2"
)

func NotificationRoute(app *fiber.App, handler handlers.NotificationsHandler) {

	notification := app.Group("/api/notification")

	/*Get a notification*/
	notification.Get("/page=:page", auth.JWTProtected(), handler.GetNotifications)
	notification.Get("/id=:id", handler.GetCompleteNotification)

}
