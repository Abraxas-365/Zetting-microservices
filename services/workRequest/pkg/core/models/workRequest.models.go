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
	Owner   interface{} `bson:"owner" json:"owner,omitempty"`
	Project interface{} `bson:"project" json:"project,omitempty"`
	Worker  interface{} `bson:"worker" json:"worker,omitempty"`
	Status  string      `bson:"status" json:"status,omitempty"`
	Created time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type WorkRequests []*WorkRequest
