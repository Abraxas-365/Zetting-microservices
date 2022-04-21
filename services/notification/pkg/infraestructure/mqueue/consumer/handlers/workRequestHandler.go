package mqHandler

import (
	"encoding/json"
	"notifications/pkg/core/models"

	"github.com/streadway/amqp"
)

type workRequest struct {
	ID      interface{} `json:"id,omitempty"`
	Owner   user        `json:"owner,omitempty"`
	Worker  user        `json:"worker,omitemptyo"`
	Project project     `json:"project,omitempty"`
}

type user struct {
	ID    interface{} `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Image string      `json:"image,omitempty"`
}

type project struct {
	ID    interface{} `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Image string      `json:"image,omitempty"`
}

func (h *mqHandler) WorkRequest(d amqp.Delivery) bool {
	workRequest := new(workRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}
	notification := models.Notification{
		NotifierUser: models.User(workRequest.Owner),
		NotifiedUser: models.User(workRequest.Worker),
		Type:         "workrequest",
		WorkRequest: models.WorkRequest{
			ID:      workRequest.ID,
			Project: models.Project(workRequest.Project),
		},
	}

	if err := h.application.CreateNotification(notification); err != nil {
		return false
	}
	return true
}
