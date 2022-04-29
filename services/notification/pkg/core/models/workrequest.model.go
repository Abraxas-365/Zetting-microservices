package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type WorkRequest struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   User      `bson:"owner,omitempty" json:"owner,omitempty"`
	Worker  User      `bson:"worker,omitempty" json:"worker,omitempty"`
	Project Project   `bson:"project" json:"project,omitempty"`
	Status  string    `bson:"status" json:"status,omitempty"`
	Created time.Time `bson:"created_at" json:"created_at,omitempty"`
}
type WorkRequestDataForNotification struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Status  string    `bson:"status" json:"status,omitempty"`
	Project Project   `bson:"project" json:"project,omitempty"`
	Created time.Time `bson:"created_at" json:"created_at,omitempty"`
}

//TODO check why the status and Created fields are empty
func (w *WorkRequest) ToNotification(tipo string) Notification {
	newNotification := Notification{
		ID:           uuid.New(),
		NotifierUser: w.Owner,
		NotifiedUser: w.Worker,
		Type:         tipo,
		WorkRequest: WorkRequestDataForNotification{
			ID:      w.ID,
			Project: w.Project,
			Status:  w.Status,
			Created: w.Created,
		},
		ReferenceId: w.ID,

		Created: time.Now(),
		Updated: time.Now(),
		Read:    false,
	}
	fmt.Println("Notification:", newNotification)
	return newNotification

}
