package ports

import (
	"projects/project/core/events"
)

type ProjectMQPublisher interface {
	PublishEvent(events.Event) error
	PublishEvents(events.Events) error
}
