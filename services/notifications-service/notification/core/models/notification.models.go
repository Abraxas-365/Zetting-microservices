package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID           uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser uuid.UUID `bson:"notified" json:"notified,omitempty"`
	NotifierUser uuid.UUID `bson:"notifier" json:"notifier,omitempty"`
	WorkRequest  uuid.UUID `bson:"workrequest " json:"workrequest,omitempty"`
	Project      uuid.UUID `bson:"project " json:"project,omitempty"`
	Type         string    `bson:"type" json:"type,omitempty"`
	Read         bool      `bson:"is_read"  json:"is_read,omitempty"`
	Message      string    `bson:"message" json:"message,omitempty"`
	Created      time.Time `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification
