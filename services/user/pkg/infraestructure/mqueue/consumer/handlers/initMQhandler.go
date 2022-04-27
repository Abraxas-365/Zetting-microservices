package mqHandler

import (
	"user/pkg/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	WorkRequest(d amqp.Delivery) bool
}
type mqHandler struct {
	userApplication application.UserApplication
}

func NewMQHandler(Service application.UserApplication) MQHandler {
	return &mqHandler{
		Service,
	}
}
