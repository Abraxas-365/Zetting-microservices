package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"projects/pkg/core/events"
)

func (mq *mqPublisher) PublishEvent(event events.Event) error {
	fmt.Println(event.Name())
	eventBytes := new(bytes.Buffer)
	json.NewEncoder(eventBytes).Encode(event)
	if err := mq.publisher.PublishToExange(event.Exchange(), event.Routing(), eventBytes.Bytes()); err != nil {
		fmt.Println("Rabbit consumer closed - critical Error")
		os.Exit(1)
		fmt.Println(err.Error())
		return err
	}
	return nil
}
