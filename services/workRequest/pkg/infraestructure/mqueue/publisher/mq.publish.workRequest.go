package mqpublisher

import (
	"bytes"
	"encoding/json"
	"work-request/pkg/core/models"
)

func (mq *mqPublisher) PublishNewWorkRequest(workRequest models.WorkRequest, exchange string, routingKey string) error {

	workRequestBytes := new(bytes.Buffer)
	type notification struct {
		NotifiedUser interface{} `json:"notified,omitempty"`
		NotifierUser interface{} `json:"notifier,omitempty"`
		Reference    interface{} `json:"reference,omitempty"`
	}
	notificationSend := notification{
		NotifiedUser: workRequest.Worker,
		NotifierUser: workRequest.Owner,
		Reference:    workRequest.ID,
	}
	json.NewEncoder(workRequestBytes).Encode(notificationSend)
	if err := mq.publisher.PublishToExange(exchange, routingKey, workRequestBytes.Bytes()); err != nil {
		return err
	}
	return nil
}
