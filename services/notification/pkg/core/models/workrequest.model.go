package models

import "time"

type WorkRequest struct {
	ID      interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   User        `json:"owner,omitempty"`
	Worker  User        `json:"worker,omitemptyo"`
	Project Project     `bson:"project" json:"project,omitempty"`
	Created time.Time   `bson:"created_at" json:"created_at,omitempty"`
}

func (w *WorkRequest) ToNotification() Notification {
	return Notification{
		NotifierUser: w.Owner,
		NotifiedUser: w.Worker,
		Type:         "workrequest",
		WorkRequest: WorkRequest{
			ID:      w.ID,
			Project: w.Project,
		},
	}

}
