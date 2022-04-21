package models

import (
	"errors"
	"time"
)

//TODO validate status
//Status{
// "P"=pending
// "A"=accepted
// "D"=denied
// }

var (
	ErrInvalidStatus = errors.New("Invalid status")
)

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

func (wr *WorkRequest) Validate() error {
	switch wr.Status {
	case "A":
	case "D":
	default:
		return ErrInvalidStatus
	}
	if err := wr.Project.Validate(); err != nil {
		return err
	}
	if err := wr.Owner.Validate(); err != nil {
		return err
	}
	if err := wr.Worker.Validate(); err != nil {
		return err
	}
	return nil
}

func (wr *WorkRequest) AnserWorkrequest() {
	wr = &WorkRequest{
		ID:      wr.ID,
		Status:  wr.Status,
		Updated: wr.Updated,
	}
}
