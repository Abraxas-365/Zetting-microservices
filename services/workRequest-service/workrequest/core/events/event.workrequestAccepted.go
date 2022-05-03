package events

import "github.com/google/uuid"

type WorkrequestAccepted struct {
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   uuid.UUID `bson:"owner" json:"owner,omitempty"`
	Worker  uuid.UUID `bson:"worker" json:"worker,omitemptyo"`
	Project uuid.UUID `bson:"project" json:"project,omitempty"`
}

func (w WorkrequestAccepted) Name() string {
	return "event.workrequest.accepted"

}
func (w WorkrequestAccepted) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestAccepted) Routing() string {
	return "accepted"
}
