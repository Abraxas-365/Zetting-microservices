package ports

import (
	"notifications/notification/core/models"

	"github.com/google/uuid"
)

type NotificationRepository interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId uuid.UUID, page int) (models.LookupNotifications, error)
	GetNotificationById(notificationId uuid.UUID) (models.LookupNotification, error)
}
