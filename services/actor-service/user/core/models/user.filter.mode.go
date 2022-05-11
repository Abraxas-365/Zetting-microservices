package models

type Filter struct {
	Gender   string   `bson:"gender" json:"gender,omitempty"`
	AgeMax   int      `bson:"age_max" json:"age_max,omitempty"`
	AgeMin   int      `bson:"age_min" json:"age_min,omitempty"`
	Features Features `bson:"features" json:"features,omitempty"`
}
