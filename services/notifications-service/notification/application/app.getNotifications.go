package application

import (
	"notifications/notification/core/models"

	"github.com/google/uuid"
)

func (s *notificationApplication) GetNotifications(userId uuid.UUID, page int) (models.LookupNotifications, error) {
	notifications, err := s.notificationRepo.GetNotifications(userId, page)
	if err != nil || notifications == nil {
		return models.LookupNotifications{}, nil
	}
	return notifications, nil

}
