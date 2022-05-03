package mqHandler

import (
	"encoding/json"
	"notifications/workrequest/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) WorkRequestCreated(d amqp.Delivery) bool {
	workRequest := new(models.WorkRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	if err := h.application.CreateWorkrequestNotification(*workRequest); err != nil {
		return false
	}

	return true
}
