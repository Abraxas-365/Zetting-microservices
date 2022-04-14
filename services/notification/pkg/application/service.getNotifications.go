package service

import "notifications/pkg/core/models"

func (s *notificationService) GetNotifications(userId interface{}, page int) (models.Notifications, error) {
	notifications, err := s.notificationRepo.GetNotifications(userId, page)
	if err != nil || notifications == nil {
		return []*models.Notification{}, nil
	}
	return notifications, nil

}
