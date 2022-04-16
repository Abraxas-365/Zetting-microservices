package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"notifications/pkg/core/models"
)

func (s *notificationService) GetCompleteNotification(notificationId interface{}) (*models.Notification, error) {
	//get the notification from the db
	notification, err := s.notificationRepo.GetNotificationById(notificationId)
	if err != nil {
		return nil, err
	}
	user, err := getUserInfo(notification.NotifierUser.(string))
	if err != nil {
	}
	//TODO: swithc for evey type of notification
	workRequest, err := getWorkRequestInfo(notification.Reference.(string))
	if err != nil {
	}

	notification.NotifierUser = user
	notification.Reference = workRequest

	return notification, nil
}

func getUserInfo(userId string) (interface{}, error) {
	user := struct {
		Name        string `json:"name,omitempty"`
		PerfilImage string `json:"perfil_image,omitempty"`
	}{}

	//go func to call the microservices of user,project
	responseUser, err := http.Get("http://zetting-user:3000/api/users/id=" + userId)
	if err != nil {
		return nil, err
	}
	responseUserData, err := ioutil.ReadAll(responseUser.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(responseUserData, &user)
	return user, nil
}
func getWorkRequestInfo(workRequestId string) (interface{}, error) {
	workRequest := struct {
		ProjectId interface{} `json:"project,omitempty"`
		Status    string      `json:"status,omitempty"`
	}{}

	//go func to call the microservices of workRequest,workRequest
	responseWorkRequest, err := http.Get("http://zetting-workrequest:3003/api/work-request/id=" + workRequestId)
	if err != nil {
		return nil, err
	}
	responseWorkRequestData, err := ioutil.ReadAll(responseWorkRequest.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(responseWorkRequestData, &workRequest)
	return workRequest, nil
}
