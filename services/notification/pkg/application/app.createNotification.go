package application

import (
	"notifications/pkg/core/models"
)

func (s *notificationApplication) CreateNotification(newNotification models.Notification) error {

	if err := s.notificationService.CreateNotification(newNotification); err != nil {
		return err
	}
	return nil

}
