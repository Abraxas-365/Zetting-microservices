package application

import (
	"fmt"
	notification "notifications/notification/core/models"
	"notifications/workrequest/core/models"
)

func (a *notificationApplication) CreateWorkrequestNotification(workrequest models.WorkRequest) error {
	newNotification := notification.Notification{}
	switch {
	//TODO Change notification.message hardcoded to the db
	case workrequest.Status == "P":
		newNotification = workrequest.ToNotification("Work request from")
	case workrequest.Status == "A":
		newNotification = workrequest.ToNotification("Is now part of your team")
	}
	if err := a.notificationApp.CreateNotification(newNotification); err != nil {
		fmt.Printf("Error creating notification")

		return err

	}
	return nil

}
