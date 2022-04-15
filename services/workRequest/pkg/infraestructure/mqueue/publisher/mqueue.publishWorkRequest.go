package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"work-request/pkg/core/models"

	"github.com/streadway/amqp"
)

func (mq *mQueue) Publish(workRequest models.WorkRequest) error {
	conn, err := amqp.Dial(mq.uri)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		mq.channelName, // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	workRequestBytes := new(bytes.Buffer)
	type notification struct {
		NotifiedUserId interface{} `json:"notified_id,omitempty"`
		NotifierUserId interface{} `json:"notifier_id,omitempty"`
		ReferenceId    interface{} `json:"reference_id,omitempty"`
	}
	notificationSend := notification{
		NotifiedUserId: workRequest.WorkerId,
		NotifierUserId: workRequest.OwnerId,
		ReferenceId:    workRequest.ID,
	}
	json.NewEncoder(workRequestBytes).Encode(notificationSend)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(workRequestBytes.Bytes()),
		})
	return nil
}
