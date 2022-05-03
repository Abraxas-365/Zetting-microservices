package application

import (
	"notifications/notification/core/models"

	"github.com/google/uuid"
)

func (s *notificationApplication) GetCompleteNotification(notificationId uuid.UUID) (models.LookupNotification, error) {
	//get the notification from the db
	notification, err := s.notificationRepo.GetNotificationById(notificationId)
	if err != nil {
		return models.LookupNotification{}, err
	}
	return notification, nil
}
