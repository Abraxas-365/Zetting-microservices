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
