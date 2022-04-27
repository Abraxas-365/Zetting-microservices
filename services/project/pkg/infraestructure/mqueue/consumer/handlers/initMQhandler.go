package mqHandler

import (
	"projects/pkg/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	AnsweredWorkRequest(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.ProjectApplication
}

func NewMQHandler(application application.ProjectApplication) MQHandler {
	return &mqHandler{
		application,
	}
}
