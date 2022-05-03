package mqHandler

import (
	"encoding/json"
	"projects/workrequest/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) WorkRequestAccepted(d amqp.Delivery) bool {

	workRequest := new(models.WorkRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	if err := h.application.WorkRequestAccepted(*workRequest); err != nil {
		return false
	}
	return true
}
