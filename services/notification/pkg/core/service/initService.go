package service

import (
	"errors"
	"notifications/pkg/core/models"
	"notifications/pkg/core/ports"
)

var (
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrEmptyParams        = errors.New("Empty parameters")
	ErrNotificationExists = errors.New("Notification exists")
)

type NotificationService interface {
	CreateNotification(newNotification models.Notification) error
}
type notificationService struct {
	notificationRepo ports.NotificationRepository
}

func NewNotificationService(notificationRepo ports.NotificationRepository) NotificationService {
	return &notificationService{
		notificationRepo,
	}
}
