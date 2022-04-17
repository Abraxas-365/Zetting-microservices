package mqHandler

import (
	"notifications/pkg/core/ports"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	WorkRequest(d amqp.Delivery) bool
}
type mqHandler struct {
	service ports.NotificationService
}

func NewMQHandler(notificationService ports.NotificationService) MQHandler {
	return &mqHandler{
		notificationService,
	}
}
