package models

import "time"

type Notification struct {
	ID           interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	NotifiedUser interface{} `bson:"notified" json:"notified,omitempty"`
	NotifierUser interface{} `bson:"notifier" json:"notifier,omitempty"`
	Reference    interface{} `bson:"reference" json:"reference,omitempty"`
	Type         string      `bson:"type" json:"type,omitempty"`
	Read         bool        `bson:"is_read"  json:"is_read,omitempty"`
	Message      string      `bson:"message" json:"message,omitempty"`
	Created      time.Time   `bson:"created_at" json:"created_at,omitempty"`
	Updated      time.Time   `bson:"updated_at" json:"updated_at,omitempty"`
}

type Notifications []*Notification
