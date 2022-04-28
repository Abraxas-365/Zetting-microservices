package ports

import (
	"work-request/pkg/core/events"
)

type WorkRequestMQPublisher interface {
	PublishEvent(events.Event) error
	PublishEvents(events.Events) error
}
