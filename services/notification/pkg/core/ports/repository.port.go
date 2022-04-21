package ports

import "notifications/pkg/core/models"

type NotificationRepository interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId interface{}, page int) (models.Notifications, error)
	GetNotificationById(notificationId interface{}) (models.Notification, error)
}
