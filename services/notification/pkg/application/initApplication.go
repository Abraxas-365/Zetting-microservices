package application

import (
	"errors"
	"notifications/pkg/core/models"
	"notifications/pkg/core/ports"
	"notifications/pkg/core/service"

	"github.com/google/uuid"
)

var (
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrUserNotFound       = errors.New("User not found")
	ErrNotificationExists = errors.New("Notification exists")
)

type NotificationApplication interface {
	CreateNotification(newNotification models.Notification) error
	GetNotifications(userId uuid.UUID, page int) (models.Notifications, error)
	GetCompleteNotification(notificationId uuid.UUID) (models.Notification, error)
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
