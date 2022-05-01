package models

import (
	"fmt"
	"notifications/notification/core/models"
	"time"

	"github.com/google/uuid"
)

type WorkRequest struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   uuid.UUID `bson:"owner,omitempty" json:"owner,omitempty"`
	Worker  uuid.UUID `bson:"worker,omitempty" json:"worker,omitempty"`
	Project uuid.UUID `bson:"project" json:"project,omitempty"`
	Status  string    `bson:"status" json:"status,omitempty"`
	Created time.Time `bson:"created_at" json:"created_at,omitempty"`
}

//TODO check why the status and Created fields are empty
func (w *WorkRequest) ToNotification(msg string) models.Notification {
	newNotification := models.Notification{
		ID:           uuid.New(),
		NotifierUser: w.Owner,
		NotifiedUser: w.Worker,
		WorkRequest:  w.ID,
		Project:      w.Project,
		Created:      w.Created,
		Updated:      time.Now(),
		Read:         false,
		Message:      msg,
	}
	fmt.Println("Notification:", newNotification)
	return newNotification

}
