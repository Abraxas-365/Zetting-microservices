package mqHandler

import (
	"encoding/json"
	"work-request/project/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) ProjectCreated(d amqp.Delivery) bool {

	project := new(models.Project)
	if err := json.Unmarshal([]byte(d.Body), &project); err != nil {
		return false
	}

	if err := h.application.CreateProject(*project); err != nil {
		return false
	}
	return true
}
