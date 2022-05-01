package events

import (
	"github.com/google/uuid"
)

type WorkrequestCreated struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   uuid.UUID `bson:"owner" json:"owner,omitempty"`
	Worker  uuid.UUID `bson:"worker" json:"worker,omitemptyo"`
	Project uuid.UUID `bson:"project" json:"project,omitempty"`
	Status  string    `bson:"status" json:"status,omitempty"`
}

func (w WorkrequestCreated) Name() string {
	return "event.workrequest.created"

}
func (w WorkrequestCreated) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestCreated) Routing() string {
	return "new"
}
