package mqHandler

import (
	"projects/workrequest/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	WorkRequestAccepted(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.WorkRequestApplication
}

func NewMQHandler(application application.WorkRequestApplication) MQHandler {
	return &mqHandler{
		application,
	}
}
