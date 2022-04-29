package events

import "work-request/pkg/core/models"

type WorkrequestCreated struct {
	ID      interface{}    `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   models.User    `bson:"owner" json:"owner,omitempty"`
	Worker  models.User    `bson:"worker" json:"worker,omitemptyo"`
	Project models.Project `bson:"project" json:"project,omitempty"`
	Status  string         `bson:"status" json:"status,omitempty"`
}

func (w WorkrequestCreated) Name() string {
	return "event.workrequest.created"

}
func (w WorkrequestCreated) Exchange() string {
	return "WorkRequest"
}
func (w WorkrequestCreated) Routing() string {
	return "workrequest_new"
}
