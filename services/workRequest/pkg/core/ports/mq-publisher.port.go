package ports

import "work-request/pkg/core/models"

type WorkRequestMQPublisher interface {
	NewWorkRequest(workRequest models.WorkRequest, exchange string, routingKey string) error
	AnswerWorkRequest(answerWorkrequest models.WorkRequest, exchange string, routingKey string) error
}
