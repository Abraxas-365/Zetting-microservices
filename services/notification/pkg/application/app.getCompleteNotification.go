package application

import (
	"notifications/pkg/core/models"
)

func (s *notificationApplication) GetCompleteNotification(notificationId interface{}) (models.Notification, error) {
	//get the notification from the db
	notification, err := s.notificationRepo.GetNotificationById(notificationId)
	if err != nil {
		return models.Notification{}, err
	}
	return notification, nil
}
