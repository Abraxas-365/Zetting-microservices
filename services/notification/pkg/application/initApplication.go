package application

import (
	"errors"
	"notifications/pkg/core/models"
	"notifications/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type NotificationApplication interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId interface{}, page int) (models.Notifications, error)
	GetCompleteNotification(notificationId interface{}) (models.Notification, error)
}
type notificationApplication struct {
	notificationRepo ports.NotificationRepository
}

func NewNotificationApplication(notificationRepo ports.NotificationRepository) NotificationApplication {
	return &notificationApplication{
		notificationRepo,
	}

}
