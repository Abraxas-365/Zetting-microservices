package mqHandler

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type workRequest struct {
	Owner  user `json:"owner,omitempty"`
	Worker user `json:"worker,omitemptyo"`
}

type user struct {
	ID    interface{} `json:"id,omitempty"`
	Name  string      `json:"name,omitempty"`
	Image string      `json:"image,omitempty"`
}

func (h *mqHandler) WorkRequest(d amqp.Delivery) bool {
	workRequest := new(workRequest)
	if err := json.Unmarshal([]byte(d.Body), &workRequest); err != nil {
		return false
	}

	return true
}
