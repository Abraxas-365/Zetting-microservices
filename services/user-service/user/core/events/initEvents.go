package events

type Event interface {
	Name() string
	Exchange() string
	Routing() string
}
type Events []Event
