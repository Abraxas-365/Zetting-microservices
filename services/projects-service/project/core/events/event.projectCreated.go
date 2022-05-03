package events

import "github.com/google/uuid"

type ProjectCreated struct {
	ID          uuid.UUID `bson:"_id,omitempty" json:"id"`
	ProjectName string    `bson:"name" json:"name,omitempty"`
	Image       string    `bson:"image" json:"image,omitempty"`
	Color       string    `bson:"color" json:"color,omitempty"`
}

func (w ProjectCreated) Name() string {
	return "event.project.created"

}
func (w ProjectCreated) Exchange() string {
	return "Project"
}
func (w ProjectCreated) Routing() string {
	return "created"
}
