package ports

import (
	"work-request/pkg/core/events"
)

type WorkRequestMQPublisher interface {
	NewWorkRequest(event events.Event) error
	AnswerWorkRequest(event events.Event) error
}
