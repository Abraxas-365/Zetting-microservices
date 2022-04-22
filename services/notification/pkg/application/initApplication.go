package application

import (
	"errors"
	"notifications/pkg/core/models"
	"notifications/pkg/core/ports"
	"notifications/pkg/core/service"
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
	notificationRepo    ports.NotificationRepository
	notificationService service.NotificationService
}

func NewNotificationApplication(notificationRepo ports.NotificationRepository, notificationService service.NotificationService) NotificationApplication {
	return &notificationApplication{
		notificationRepo,
		notificationService,
	}

}