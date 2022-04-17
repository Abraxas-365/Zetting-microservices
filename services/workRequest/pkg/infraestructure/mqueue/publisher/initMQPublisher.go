package mqpublisher

import (
	"work-request/internal/rabbit"
	"work-request/pkg/core/ports"
)

type mqPublisher struct {
	publisher rabbit.MQueue
}

func NewMQPublisher(mq rabbit.MQueue) ports.WorkRequestMQPublisher {
	return &mqPublisher{
		mq,
	}
}
