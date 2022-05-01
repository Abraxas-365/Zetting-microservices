package mqHandler

import (
	"notifications/workrequest/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	WorkRequestCreated(d amqp.Delivery) bool
	WorkRequestAccepted(d amqp.Delivery) bool
	WorkRequestDenied(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.WorkRequestApplication
}

func NewMQHandler(workRequestApplication application.WorkRequestApplication) MQHandler {
	return &mqHandler{
		workRequestApplication,
	}
}
