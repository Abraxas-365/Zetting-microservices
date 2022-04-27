package events

import "work-request/pkg/core/models"

type WorkrequestDenied struct {
	ID      interface{}    `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   models.User    `bson:"owner" json:"owner,omitempty"`
	Worker  models.User    `bson:"worker" json:"worker,omitemptyo"`
	Project models.Project `bson:"project" json:"project,omitempty"`
}

func (w WorkrequestDenied) Name() string {
	return "event.workrequest.denied"

}
func (w WorkrequestDenied) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestDenied) Routing() string {
	return "workrequest_denied"
}
