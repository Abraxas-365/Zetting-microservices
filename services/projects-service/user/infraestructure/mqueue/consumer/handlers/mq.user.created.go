package mqHandler

import (
	"encoding/json"
	"fmt"
	"projects/user/core/models"

	"github.com/streadway/amqp"
)

func (h *mqHandler) UserCreated(d amqp.Delivery) bool {
	fmt.Println("USER CREATED")

	user := new(models.User)
	if err := json.Unmarshal([]byte(d.Body), &user); err != nil {
		return false
	}
	fmt.Println(user)

	if err := h.application.CreateUser(*user); err != nil {
		return false
	}
	return true
}
