package models

import (
	"time"
)

//TODO validate status
//Status{
// "P"=pending
// "A"=accepted
// "D"=denied
// }
type WorkRequest struct {
	ID      interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   User        `bson:"owner" json:"owner,omitempty"`
	Worker  User        `bson:"worker" json:"worker,omitemptyo"`
	Project Project     `bson:"project" json:"project,omitempty"`
	Status  string      `bson:"status" json:"status,omitempty"`
	Created time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type WorkRequests []*WorkRequest

type User struct {
	ID    interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"image" json:"image,omitempty"`
}

type Project struct {
	ID    interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string      `bson:"name" json:"name,omitempty"`
	Image string      `bson:"image" json:"image,omitempty"`
}
