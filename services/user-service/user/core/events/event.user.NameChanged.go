package events

import "github.com/google/uuid"

type UserNameChanged struct {
	ID       uuid.UUID `bson:"_id,omitempty" json:"id"`
	UserName string    `bson:"name" json:"name,omitempty"`
}

func (e UserNameChanged) Name() string {
	return "event.user.name_changed"
}

func (e UserNameChanged) Exchange() string {
	return "User"
}
func (e UserNameChanged) Routing() string {
	return "name"
}
