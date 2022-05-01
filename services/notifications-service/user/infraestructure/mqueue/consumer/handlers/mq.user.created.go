package mqHandler

import (
	"encoding/json"
	"notifications/user/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) UserCreated(d amqp.Delivery) bool {

	user := new(models.User)
	if err := json.Unmarshal([]byte(d.Body), &user); err != nil {
		return false
	}

	if err := h.application.CreateUser(*user); err != nil {
		return false
	}
	return true
}
