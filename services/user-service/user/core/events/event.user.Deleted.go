package events

import "github.com/google/uuid"

type UserDeleted struct {
	ID uuid.UUID `bson:"_id,omitempty" json:"id"`
}

func (e UserDeleted) Name() string {
	return "event.user.deleted"
}

func (e UserDeleted) Exchange() string {
	return "User"
}
func (e UserDeleted) Routing() string {
	return "deleted"
}
