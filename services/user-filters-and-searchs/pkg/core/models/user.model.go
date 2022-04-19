package models

import (
	"time"
)

type User struct {
	ID          interface{} `bson:"_id,omitempty" json:"id"`
	Name        string      `bson:"name" json:"name,omitempty"`
	Password    string      `bson:"password" json:"password,omitempty"`
	PerfilImage string      `bson:"perfil_image" json:"perfil_image,omitempty"`
	Contact     Contact     `bson:"contact" json:"contact,omitempty"`
	Profession  Profession  `bson:"profession" json:"profession,omitempty"`
	Features    Features    `bson:"features" json:"features,omitempty"`
	Projects    Projects    `bson:"projects" json:"projects,omitempty"`
	Verified    bool        `bson:"verified" json:"verified,omitempty"`
	Created     time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

//To get users wirth out the password
type GetUsers struct {
	ID          interface{} `bson:"_id,omitempty" json:"id"`
	Name        string      `bson:"name" json:"name,omitempty"`
	PerfilImage string      `bson:"perfil_image" json:"perfil_image,omitempty"`
	Contact     Contact     `bson:"contact" json:"contact,omitempty"`
	Profession  Profession  `bson:"profession" json:"profession,omitempty"`
	Features    Features    `bson:"features" json:"features,omitempty"`
	Verified    bool        `bson:"verified" json:"verified,omitempty"`
	Created     time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated     time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}
type Users []*GetUsers
