package ports

import (
	"work-request/workrequest/core/events"
)

type WorkRequestMQPublisher interface {
	PublishEvent(events.Event) error
	PublishEvents(events.Events) error
}
