package service

import (
	"notifications/pkg/core/models"
)

func (s *notificationService) GetCompleteNotification(notificationId interface{}) (*models.Notification, error) {
	//get the notification from the db
	notification, err := s.notificationRepo.GetNotificationById(notificationId)
	if err != nil {
		return nil, err
	}

	return notification, nil
}
