package mqHandler

import (
	"github.com/streadway/amqp"
	"user/pkg/core/ports"
)

type MQHandler interface {
	WorkRequest(d amqp.Delivery) bool
}
type mqHandler struct {
	userApplication ports.UserApplication
}

func NewMQHandler(Service ports.UserApplication) MQHandler {
	return &mqHandler{
		Service,
	}
}
