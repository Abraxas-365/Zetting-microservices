package mqHandler

import (
	"encoding/json"
	"notifications/pkg/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) WorkRequestDenied(d amqp.Delivery) bool {
	workRequest := new(models.WorkRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	if err := h.application.CreateNotification(workRequest.ToNotification("workrequest_denied")); err != nil {
		return false
	}
	return true
}
