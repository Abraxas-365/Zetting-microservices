package service

import (
	"errors"
	"notifications/pkg/core/ports"
)

var (
	ErrUnauthorized = errors.New("Unauthorized")
	ErrUserNotFound = errors.New("User not found")
)

type notificationService struct {
	notificationRepo ports.NotificationRepository
}

func NewNotificationService(notificationRepo ports.NotificationRepository) ports.NotificationService {
	return &notificationService{
		notificationRepo,
	}

}
