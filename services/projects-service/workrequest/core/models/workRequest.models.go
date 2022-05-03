package models

import "github.com/google/uuid"

type WorkRequest struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Worker  uuid.UUID `json:"worker,omitempty"`
	Project uuid.UUID `json:"project,omitempty"`
}
