package application

import (
	"errors"
	"notifications/notification/application"
	"notifications/workrequest/core/models"
)

var (
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrUserNotFound       = errors.New("User not found")
	ErrNotificationExists = errors.New("Notification exists")
)

type WorkRequestApplication interface {
	CreateWorkrequestNotification(workrequest models.WorkRequest) error
}
type notificationApplication struct {
	notificationApp application.NotificationApplication
}

func NewWorkRequestApplication(notificationApp application.NotificationApplication) WorkRequestApplication {
	return &notificationApplication{
		notificationApp,
	}

}
