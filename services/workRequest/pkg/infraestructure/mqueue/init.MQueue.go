package mqueue

import (
	"errors"
	"work-request/pkg/core/ports"

	"github.com/streadway/amqp"
)

var (
	ErrConflict = errors.New("Allready exist")
)

type mQueue struct {
	uri         string
	channelName string
	channel     *amqp.Channel

	// conn    *amqp.Connection
}

func NewMQueue(uri string, channelName string) (ports.WorkRequestMQueue, error) {
	mqueue := &mQueue{
		channelName: channelName,
		uri:         uri,
	}

	return mqueue, nil

}
