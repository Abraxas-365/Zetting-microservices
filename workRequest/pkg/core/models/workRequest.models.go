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
	ID        interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	OwnerId   interface{} `bson:"owner_id" json:"owner_id,omitempty"`
	ProjectId interface{} `bson:"project_id" json:"project_id,omitempty"`
	WorkerId  interface{} `bson:"worker_id" json:"worker_id,omitempty"`
	Status    string      `bson:"status" json:"status,omitempty"`
	Created   time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated   time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type WorkRequests []*WorkRequest
