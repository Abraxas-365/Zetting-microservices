package mqconsumer

import (
	"encoding/json"
	"fmt"
	"notifications/pkg/core/models"
)

func (mq *mQueue) ConsumeExchange(exchange string, routingKey string) error {
	err := mq.channel.ExchangeDeclare(
		exchange,
		routingKey,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	q, err := mq.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	err = mq.channel.QueueBind(
		q.Name,        // queue name
		routingKey,    // routing key
		"logs_direct", // exchange
		false,
		nil,
	)
	msgs, err := mq.channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)

	for msg := range msgs {
		workRequest := new(models.Notification)
		json.Unmarshal(msg.Body, workRequest)
		workRequest.Type = mq.channelName
		fmt.Println(workRequest)
		mq.service.CreateNotification(*workRequest)
	}

	return nil

}
