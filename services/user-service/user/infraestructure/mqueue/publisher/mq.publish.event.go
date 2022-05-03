package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"user/user/core/events"
)

func (mq *mqPublisher) PublishEvent(event events.Event) error {
	fmt.Println(event.Name())
	fmt.Println(event)

	eventBytes := new(bytes.Buffer)
	json.NewEncoder(eventBytes).Encode(event)
	if err := mq.publisher.PublishToExchange(event.Exchange(), event.Routing(), eventBytes.Bytes()); err != nil {
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
		return err
	}

	return nil
}
