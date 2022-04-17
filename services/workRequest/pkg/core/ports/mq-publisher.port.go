package ports

import "work-request/pkg/core/models"

type WorkRequestMQPublisher interface {
	PublishNewWorkRequest(workRequest models.WorkRequest, exchange string, routingKey string) error
}
