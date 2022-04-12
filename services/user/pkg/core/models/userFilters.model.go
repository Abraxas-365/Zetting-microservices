package models

type Filter struct {
	AgeTop int `json:"age_top"`
	AgeLow int `json:"age_low"`
	// desde aqui es actores
	HeightTop  int      `json:"height_top"`
	HeightLow  int      `json:"height_low"`
	Body       string   `json:"body"`
	Skin       string   `json:"skin"`
	HairType   string   `json:"hair_type"`
	HairZise   string   `json:"hair_zise"`
	Skills     []string `json:"skills"`
	FacialHair string   `json:"facial_hair"`
	HairColor  string   `json:"hair_color"`
	EyeColor   string   `json:"eye_color"`
}
