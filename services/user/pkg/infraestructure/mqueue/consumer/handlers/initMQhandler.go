package mqHandler

import (
	"github.com/streadway/amqp"
	"user/pkg/core/ports"
)

type MQHandler interface {
	WorkRequest(d amqp.Delivery) bool
}
type mqHandler struct {
	service ports.UserService
}

func NewMQHandler(Service ports.UserService) MQHandler {
	return &mqHandler{
		notificationService,
	}
}
