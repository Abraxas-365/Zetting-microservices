package models

/*Profession obj*/
type Profession string

func (p Profession) Validate() bool {
	if p == "actor" {
		return true
	}
	return false
}

/*Features obj*/
type Features struct {
	Height     int      `bson:"height" json:"height,omitempty"`
	Body       string   `bson:"body" json:"body,omitempty"`
	Skin       string   `bson:"skin" json:"skin,omitempty"`
	HairType   string   `bson:"hair_type" json:"hair_type,omitempty"`
	HairZise   string   `bson:"hair_zise" json:"hair_zise,omitempty"`
	Skills     []string `bson:"skills" json:"skills,omitempty"`
	FacialHair string   `bson:"facial_hair" json:"facial_hair,omitempty"`
	HairColor  string   `bson:"hair_color" json:"hair_color,omitempty"`
	EyeColor   string   `bson:"eye_color" json:"eye_color,omitempty"`
}
