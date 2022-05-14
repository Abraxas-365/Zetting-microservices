package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Name        string     `bson:"name" json:"name,omitempty"`
	PerfilImage string     `bson:"perfil_image" json:"perfil_image,omitempty"`
	Gender      string     `bson:"gender" json:"gender,omitempty"`
	Age         int        `bson:"age" json:"age,omitempty"`
	Features    Features   `bson:"features" json:"features,omitempty"`
	Description string     `bson:"description" json:"description"`
	Profession  Profession `bson:"profession" json:"profession,omitempty"`
}

type Users []*User
