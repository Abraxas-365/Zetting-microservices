package mqueue

import (
	"errors"
	"fmt"
	"notifications/pkg/core/ports"

	"github.com/streadway/amqp"
)

var (
	ErrConflict            = errors.New("Allready exist")
	ErrCantConnectToRabbit = errors.New("CANT CONNECT TO Rabbit")
)

type mQueue struct {
	uri         string
	channelName string
	channel     *amqp.Channel
	service     ports.NotificationService
	conn        *amqp.Connection
}

func NewMQueue(uri string, channelName string, service ports.NotificationService) (ports.NotificationMQueue, error) {
	mqueue := &mQueue{
		channelName: channelName,
		uri:         uri,
		service:     service,
	}
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, ErrCantConnectToRabbit
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()
	mqueue.conn = conn
	_, err = ch.QueueDeclare(
		channelName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	fmt.Println("STARTED CORRECTLY")
	return mqueue, nil

}
