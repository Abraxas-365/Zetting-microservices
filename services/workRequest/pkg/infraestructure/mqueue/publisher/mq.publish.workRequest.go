package mqpublisher

import (
	"bytes"
	"encoding/json"
	"work-request/pkg/core/models"
)

func (mq *mqPublisher) PublishNewWorkRequest(workRequest models.WorkRequest, exchange string, routingKey string) error {

	workRequestBytes := new(bytes.Buffer)
	json.NewEncoder(workRequestBytes).Encode(workRequest)
	if err := mq.publisher.PublishToExange(exchange, routingKey, workRequestBytes.Bytes()); err != nil {
		return err
	}
	return nil
}
