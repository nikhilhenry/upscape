package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Actionable struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty" validate:"required,max=16,alpha"`
	Oid       int                  `json:"id,omitempty" bson:"id,omitempty"`
	Inbox     bool                 `json:"inbox,omitempty" bson:"inbox,omitempty" validate:"required,bool"`
	TaskIDs   []primitive.ObjectID `json:"taskIDs,omitempty" bson:"taskIDs,omitempty"`
	Completed bool                 `json:"completed" bson:"completed" validate:"bool"`
}
