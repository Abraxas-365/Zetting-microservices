package models

type Features struct {
	Gender string `json:"gender,omitempty"`
	Age    int    `json:"age,omitempty"`
	// desde aqui es actores
	Height     int      `json:"height,omitempty"`
	Body       string   `json:"body,omitempty"`
	Skin       string   `json:"skin,omitempty"`
	HairType   string   `json:"hair_type,omitempty"`
	HairZise   string   `json:"hair_zise,omitempty"`
	Skills     []string `json:"skills,omitempty"`
	FacialHair string   `json:"facial_hair,omitempty"`
	HairColor  string   `json:"hair_color,omitempty"`
	EyeColor   string   `json:"eye_color,omitempty"`
	// Aqui acaba actores
}
