package mqpublisher

import (
	"user/internal/rabbit"
	"user/user/core/ports"
)

type mqPublisher struct {
	publisher rabbit.MQueue
}

func NewMQPublisher(mq rabbit.MQueue) ports.UserMQPublisher {
	return &mqPublisher{
		mq,
	}
}
