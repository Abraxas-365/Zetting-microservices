package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID  `bson:"_id,omitempty" json:"id"`
	Name        string     `bson:"name" json:"name,omitempty"`
	PerfilImage string     `bson:"perfil_image" json:"perfil_image,omitempty"`
	Gender      string     `json:"gender,omitempty"`
	Age         int        `json:"age,omitempty"`
	Features    Features   `bson:"features" json:"features,omitempty"`
	Description string     `bson:"description" json:"description,omitempty"`
	Profession  Profession `bson:"profession" json:"profession,omitempty"`
}

type Features struct {
	Height     int      `json:"height,omitempty"`
	Body       string   `json:"body,omitempty"`
	Skin       string   `json:"skin,omitempty"`
	HairType   string   `json:"hair_type,omitempty"`
	HairZise   string   `json:"hair_zise,omitempty"`
	Skills     []string `json:"skills,omitempty"`
	FacialHair string   `json:"facial_hair,omitempty"`
	HairColor  string   `json:"hair_color,omitempty"`
	EyeColor   string   `json:"eye_color,omitempty"`
}

type Profession string

func (p Profession) Validate() bool {
	if p == "actor" {
		return true
	}
	return false
}
