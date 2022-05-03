package events

import "github.com/google/uuid"

type UserCreated struct {
	ID          uuid.UUID `bson:"_id,omitempty" json:"id"`
	UserName    string    `bson:"name" json:"name,omitempty"`
	PerfilImage string    `bson:"perfil_image" json:"perfil_image,omitempty"`
}

func (e UserCreated) Name() string {
	return "event.user.created"
}

func (e UserCreated) Exchange() string {
	return "User"
}
func (e UserCreated) Routing() string {
	return "created"
}
