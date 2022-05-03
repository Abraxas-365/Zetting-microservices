package application

import (
	"errors"
	"github.com/google/uuid"
	"notifications/notification/core/models"
	"notifications/notification/core/ports"
)

var (
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrUserNotFound       = errors.New("User not found")
	ErrNotificationExists = errors.New("Notification exists")
)

type NotificationApplication interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId uuid.UUID, page int) (models.LookupNotifications, error)
	GetCompleteNotification(notificationId uuid.UUID) (models.LookupNotification, error)
}
type notificationApplication struct {
	notificationRepo ports.NotificationRepository
}

func NewNotificationApplication(notificationRepo ports.NotificationRepository) NotificationApplication {
	return &notificationApplication{
		notificationRepo,
	}

}
