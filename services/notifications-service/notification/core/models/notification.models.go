package models

import (
	"notifications/project/core/models"
	user "notifications/user/core/models"
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID           uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser uuid.UUID `bson:"notified" json:"notified,omitempty"`
	NotifierUser uuid.UUID `bson:"notifier" json:"notifier,omitempty"`
	WorkRequest  uuid.UUID `bson:"workrequest " json:"workrequest,omitempty"` //this will ommit if it ist a work request
	Project      uuid.UUID `bson:"project" json:"project,omitempty"`          //this will ommit if it ist a work request
	Type         string    `bson:"type" json:"type,omitempty"`
	Read         bool      `bson:"is_read"  json:"is_read,omitempty"`
	Message      string    `bson:"message" json:"message,omitempty"`
	Created      time.Time `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification

type LookupNotification struct {
	ID           uuid.UUID      `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser user.User      `bson:"notified" json:"notified,omitempty"`
	NotifierUser user.User      `bson:"notifier" json:"notifier,omitempty"`
	Project      models.Project `bson:"project" json:"project,omitempty"`          //this will ommit if it ist a work request
	WorkRequest  uuid.UUID      `bson:"workrequest " json:"workrequest,omitempty"` //this will ommit if it ist a work request
	Type         string         `bson:"type" json:"type,omitempty"`
	Read         bool           `bson:"is_read"  json:"is_read,omitempty"`
	Message      string         `bson:"message" json:"message,omitempty"`
	Created      time.Time      `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time      `bson:"updated_at" json:"updated_at,omitempty"`
}

type LookupNotifications []*LookupNotification
