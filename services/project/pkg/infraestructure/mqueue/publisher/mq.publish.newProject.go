package mqpublisher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"projects/pkg/core/models"
)

func (mq *mqPublisher) NewProject(project models.Project, exchange string, routingKey string) error {

	projectBytes := new(bytes.Buffer)
	json.NewEncoder(projectBytes).Encode(project)
	if err := mq.publisher.PublishToExange(exchange, routingKey, projectBytes.Bytes()); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
