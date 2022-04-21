package mqpublisher

import (
	"bytes"
	"encoding/json"
	"work-request/pkg/core/models"
)

func (mq *mqPublisher) AnswerWorkRequest(answerWorkrequest models.WorkRequest, exchange string, routingKey string) error {

	answerWorkrequest.AnserWorkrequest()
	answerWorkRequestBytes := new(bytes.Buffer)
	json.NewEncoder(answerWorkRequestBytes).Encode(answerWorkrequest)
	if err := mq.publisher.PublishToExange(exchange, routingKey, answerWorkRequestBytes.Bytes()); err != nil {
		return err
	}

	return nil
}
