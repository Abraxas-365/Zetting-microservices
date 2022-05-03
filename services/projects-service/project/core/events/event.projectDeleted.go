package events

type ProjectDeleted struct {
	ID interface{} `bson:"_id,omitempty" json:"id"`
}

func (w ProjectDeleted) Name() string {
	return "event.project.deleted"

}
func (w ProjectDeleted) Exchange() string {
	return "Project"
}
func (w ProjectDeleted) Routing() string {
	return "deleted"
}
