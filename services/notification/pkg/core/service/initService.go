package service

import (
	"errors"
	"notifications/pkg/core/models"
)

var (
	ErrUnauthorized       = errors.New("Unauthorized")
	ErrEmptyParams        = errors.New("Empty parameters")
	ErrNotificationExists = errors.New("Notification exists")
)

type NotificationService interface {
	CreateNotification(newNotification models.Notification) error
}

type ServiceRepository interface {
	CreateNotification(newNotification models.Notification) error
	IsNotificationExist(newNotification models.Notification) bool
}
type notificationService struct {
	notificationRepo ServiceRepository
}

func NewNotificationService(notificationRepo ServiceRepository) NotificationService {
	return &notificationService{
		notificationRepo,
	}
}
