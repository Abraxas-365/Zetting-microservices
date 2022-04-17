package mqpublisher

import (
	"projects/internal/rabbit"
	"projects/pkg/core/ports"
)

type mqPublisher struct {
	publisher rabbit.MQueue
}

func NewMQPublisher(mq rabbit.MQueue) ports.ProjectMQPublisher {
	return &mqPublisher{
		mq,
	}
}
