package events

import (
	"github.com/google/uuid"
)

type WorkrequestDenied struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   uuid.UUID `bson:"owner" json:"owner,omitempty"`
	Worker  uuid.UUID `bson:"worker" json:"worker,omitemptyo"`
	Project uuid.UUID `bson:"project" json:"project,omitempty"`
}

func (w WorkrequestDenied) Name() string {
	return "event.workrequest.denied"

}
func (w WorkrequestDenied) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestDenied) Routing() string {
	return "denied"
}
