package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"work-request/pkg/core/events"
)

func (mq *mqPublisher) PublishEvents(events events.Events) error {
	for _, e := range events {
		fmt.Println(e.Name())
		eventBytes := new(bytes.Buffer)
		json.NewEncoder(eventBytes).Encode(e)
		if err := mq.publisher.PublishToExange(e.Exchange(), e.Routing(), eventBytes.Bytes()); err != nil {
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}
