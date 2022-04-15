package mqconsumer

import (
	"encoding/json"
	"fmt"
	"notifications/pkg/core/models"
	"os"

	"github.com/streadway/amqp"
)

func (mq *mQueue) ConsumerWorkRequest() error {
	conn, err := amqp.Dial(mq.uri)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
		return ErrCantConnectToRabbit
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		mq.channelName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	for msg := range msgs {
		workRequest := new(models.Notification)
		json.Unmarshal(msg.Body, workRequest)
		workRequest.Type = mq.channelName
		mq.service.CreateNotification(*workRequest)
	}

	return nil

}
