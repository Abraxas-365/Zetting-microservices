package service

import "notifications/pkg/core/models"

func (s *notificationService) CreateNotification(newNotification models.Notification) error {
	if s.notificationRepo.IsNotificationExist(newNotification) {
		return ErrNotificationExists
	}

	switch {
	case newNotification.Type == "workrequest":
		newNotification.Message = "Work request from"

	}
	if err := s.notificationRepo.CreateNotification(newNotification); err != nil {
		return err
	}
	return nil
}
