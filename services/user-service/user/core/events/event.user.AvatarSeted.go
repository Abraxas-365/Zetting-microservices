package events

import "github.com/google/uuid"

type UserAvatarSeted struct {
	ID          uuid.UUID `bson:"_id,omitempty" json:"id"`
	PerfilImage string    `bson:"perfil_image" json:"perfil_image,omitempty"`
}

func (e UserAvatarSeted) Name() string {
	return "event.user.avatar_seted"
}

func (e UserAvatarSeted) Exchange() string {
	return "User"
}
func (e UserAvatarSeted) Routing() string {
	return "avatar"
}
