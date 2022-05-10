package mqHandler

import (
	"actor-service/user/core/models"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
)

func (h *mqHandler) UserCreated(d amqp.Delivery) bool {
	fmt.Println("USER CREATED")

	user := new(models.User)
	if err := json.Unmarshal([]byte(d.Body), &user); err != nil {
		return false
	}
	//if profession is actor safe user
	if user.Profession.Validate() {
		if err := h.application.CreateUser(*user); err != nil {
			return false
		}
	}

	return true
}
