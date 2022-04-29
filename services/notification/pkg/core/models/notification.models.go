package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID           uuid.UUID                      `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser User                           `bson:"notified" json:"notified,omitempty"`
	NotifierUser User                           `bson:"notifier" json:"notifier,omitempty"`
	ReferenceId  uuid.UUID                      `bson:"reference " json:"reference"`
	WorkRequest  WorkRequestDataForNotification `bson:"workrequest,omitempty" json:"workRequest,omitempty"`
	Type         string                         `bson:"type" json:"type,omitempty"`
	Read         bool                           `bson:"is_read"  json:"is_read,omitempty"`
	Message      string                         `bson:"message" json:"message,omitempty"`
	Created      time.Time                      `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time                      `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification

type User struct {
	ID    uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string    `bson:"name,omitempty" json:"name,omitempty"`
	Image string    `bson:"perfil_image,omitempty" json:"perfil_image,omitempty"`
}

type Project struct {
	ID    uuid.UUID `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string    `bson:"name" json:"name,omitempty"`
	Image string    `bson:"image" json:"image,omitempty"`
}
