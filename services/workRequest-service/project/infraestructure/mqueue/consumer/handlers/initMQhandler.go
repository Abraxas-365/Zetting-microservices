package mqHandler

import (
	"work-request/project/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	ProjectCreated(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.ProjectApplication
}

func NewMQHandler(application application.ProjectApplication) MQHandler {
	return &mqHandler{
		application,
	}
}
