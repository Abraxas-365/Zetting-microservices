package models

import (
	"errors"
	"time"
	project "work-request/project/core/models"
	user "work-request/user/core/models"

	"github.com/google/uuid"
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
	ID      uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   uuid.UUID `bson:"owner" json:"owner,omitempty"`
	Worker  uuid.UUID `bson:"worker" json:"worker,omitemptyo"`
	Project uuid.UUID `bson:"project" json:"project,omitempty"`
	Status  Status    `bson:"status" json:"status,omitempty"`
	Created time.Time `bson:"created_at" json:"created_at,omitempty"`
	Updated time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type LookUpWorkRequest struct {
	ID      uuid.UUID       `bson:"_id,omitempty" json:"id,omitempty"`
	Owner   user.User       `bson:"owner" json:"owner,omitempty"`
	Worker  user.User       `bson:"worker" json:"worker,omitemptyo"`
	Project project.Project `bson:"project" json:"project,omitempty"`
	Status  Status          `bson:"status" json:"status,omitempty"`
	Created time.Time       `bson:"created_at" json:"created_at,omitempty"`
	Updated time.Time       `bson:"updated_at" json:"updated_at,omitempty"`
}

type LookUpWorkRequests []*LookUpWorkRequest
type WorkRequests []*WorkRequest

func (wr *WorkRequest) New() WorkRequest {
	wr.ID = uuid.New()
	wr.Created = time.Now()
	wr.Updated = time.Now()
	wr.Status = "P"
	return *wr
}

type Status string

func (s Status) Validate() error {
	switch s {
	case "A":
	case "D":
	default:
		return ErrInvalidStatus
	}
	return nil
}
