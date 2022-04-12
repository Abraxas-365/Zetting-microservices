package ports

import "notifications/pkg/core/models"

type NotificationService interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId interface{}, page int) (models.Notifications, error)
}
