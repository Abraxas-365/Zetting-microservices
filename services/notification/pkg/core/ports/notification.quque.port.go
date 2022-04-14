package ports

type NotificationMQueue interface {
	ConsumerWorkRequest() error
}
