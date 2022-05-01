package mqHandler

import (
	"work-request/user/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	UserCreated(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.UserApplication
}

func NewMQHandler(application application.UserApplication) MQHandler {
	return &mqHandler{
		application,
	}
}
