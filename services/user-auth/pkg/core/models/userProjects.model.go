package models

type Projects struct {
	Owner  interface{} `bson:"owner" json:"owner,omitempty"`
	Worker interface{} `bson:"worker" json:"worker,omitempty"`
}
