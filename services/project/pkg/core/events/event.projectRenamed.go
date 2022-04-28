package events

type ProjectRenamed struct {
	ID interface{} `bson:"_id,omitempty" json:"id"`
}

func (w ProjectRenamed) Name() string {
	return "event.project.renamed"

}
func (w ProjectRenamed) Exchange() string {
	return "Project"
}
func (w ProjectRenamed) Routing() string {
	return "project_renamed"
}
