package application

import (
	"fmt"
	"notifications/notification/core/models"
)

func (s *notificationApplication) CreateNotification(newNotification models.Notification) error {
	if err := s.notificationRepo.CreateNotification(newNotification); err != nil {
		fmt.Printf("Error creating notification")
		return err

	}
	return nil

}
