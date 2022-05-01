package mqHandler

import (
	"encoding/json"
	"fmt"
	"notifications/workrequest/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) WorkRequestCreated(d amqp.Delivery) bool {
	workRequest := new(models.WorkRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	fmt.Println("Work request", workRequest)

	if err := h.application.CreateWorkrequestNotification(*workRequest); err != nil {
		return false
	}

	return true
}
