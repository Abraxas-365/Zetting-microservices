package mqHandler

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/streadway/amqp"
)

type workRequest struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Owner   user      `json:"owner,omitempty"`
	Worker  user      `json:"worker,omitemptyo"`
	Project project   `json:"project,omitempty"`
}

type user struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Image string    `json:"image,omitempty"`
}

type project struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Image string    `json:"image,omitempty"`
}

func (h *mqHandler) AnsweredWorkRequest(d amqp.Delivery) bool {

	workRequest := new(workRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	if err := h.application.AddUserToProject(workRequest.Worker.ID, workRequest.Project.ID, ""); err != nil {
		return false
	}
	return true
}
