package mqHandler

import (
	"notifications/pkg/application"

	"github.com/streadway/amqp"
)

type MQHandler interface {
	WorkRequestCreated(d amqp.Delivery) bool
	WorkRequestAccepted(d amqp.Delivery) bool
	WorkRequestDenied(d amqp.Delivery) bool
}
type mqHandler struct {
	application application.NotificationApplication
}

func NewMQHandler(notificationApplication application.NotificationApplication) MQHandler {
	return &mqHandler{
		notificationApplication,
	}
}
