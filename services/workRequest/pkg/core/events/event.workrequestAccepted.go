package events

import "work-request/pkg/core/models"

type WorkrequestAccepted struct {
	ID      interface{}    `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   models.User    `bson:"owner" json:"owner,omitempty"`
	Worker  models.User    `bson:"worker" json:"worker,omitemptyo"`
	Project models.Project `bson:"project" json:"project,omitempty"`
}

func (w WorkrequestAccepted) Name() string {
	return "event.workrequest.accepted"

}
func (w WorkrequestAccepted) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestAccepted) Routing() string {
	return "workrequest_accepted"
}
