package application

import (
	"notifications/pkg/core/models"

	"github.com/google/uuid"
)

func (s *notificationApplication) GetNotifications(userId uuid.UUID, page int) (models.Notifications, error) {
	notifications, err := s.notificationRepo.GetNotifications(userId, page)
	if err != nil || notifications == nil {
		return []*models.Notification{}, nil
	}
	return notifications, nil

}
