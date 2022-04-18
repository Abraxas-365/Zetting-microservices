package service

import (
	"fmt"
	"notifications/pkg/core/models"
)

func (s *notificationService) CreateNotification(newNotification models.Notification) error {

	switch {
	case newNotification.Type == "workrequest":
		newNotification.Message = "Work request from"

	}
	fmt.Println("La notificación: ", newNotification)

	if err := s.notificationRepo.CreateNotification(newNotification); err != nil {
		return err
	}
	return nil

}
