package application

import (
	"notifications/pkg/core/models"

	"github.com/google/uuid"
)

func (s *notificationApplication) GetCompleteNotification(notificationId uuid.UUID) (models.Notification, error) {
	//get the notification from the db
	notification, err := s.notificationRepo.GetNotificationById(notificationId)
	if err != nil {
		return models.Notification{}, err
	}
	return notification, nil
}
