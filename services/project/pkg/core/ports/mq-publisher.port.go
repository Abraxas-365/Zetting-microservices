package ports

import (
	"projects/pkg/core/events"
)

type ProjectMQPublisher interface {
	PublishEvent(events.Event) error
	PublishEvents(events.Events) error
}
