package mqHandler

import (
	"encoding/json"
	"fmt"
	"notifications/pkg/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) WorkRequest(d amqp.Delivery) bool {
	notification := new(models.Notification)
	if err := json.Unmarshal([]byte(d.Body), &notification); err != nil {
		return false
	}
	fmt.Println(notification)

	if err := h.service.CreateNotification(*notification); err != nil {
		return false
	}
	fmt.Println("Biennnnnnnnn")
	return true
}
