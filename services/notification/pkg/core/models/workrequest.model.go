package models

import "time"

type WorkRequest struct {
	ID      interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   User        `bson:"owner,omitempty" json:"owner,omitempty"`
	Worker  User        `bson:"worker,omitempty" json:"worker,omitempty"`
	Project Project     `bson:"project" json:"project,omitempty"`
	Created time.Time   `bson:"created_at" json:"created_at,omitempty"`
}

func (w *WorkRequest) ToNotification(tipo string) Notification {
	return Notification{
		NotifierUser: w.Owner,
		NotifiedUser: w.Worker,
		Type:         tipo,
		WorkRequest: WorkRequest{
			ID:      w.ID,
			Project: w.Project,
		},
	}

}
