package routes

import (
	"notifications/pkg/infraestructure/rest/handlers"
	"notifications/pkg/internal/auth"

	"github.com/gofiber/fiber/v2"
)

func NotificationRoute(app *fiber.App, handler handlers.NotificationsHandler) {

	notification := app.Group("/api/notification")

	/*Create a new notification*/
	notification.Post("/new/type=:type", auth.JWTProtected(), handler.CreateNotification)
	/*Get a notification*/
	notification.Get("/page=:page", auth.JWTProtected(), handler.GetNotifications)

}
