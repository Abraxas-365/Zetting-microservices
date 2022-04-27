package application

import (
	"notifications/pkg/core/models"
)

func (s *notificationApplication) CreateNotification(newNotification models.Notification) error {

	if s.notificationRepo.IsNotificationExist(newNotification) {
		return ErrNotificationExists
	}

	switch {
	//TODO Change notification.message hardcoded to the db
	case newNotification.Type == "workrequest_new":
		newNotification.Message = "Work request from"
	case newNotification.Type == "workrequest_accepted":
		newNotification.Message = "Is now part of your team"
	}
	if err := s.notificationRepo.CreateNotification(newNotification); err != nil {
		return err
	}
	return nil

}
