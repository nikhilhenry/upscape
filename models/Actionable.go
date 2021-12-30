package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Actionable struct {
	ID        primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string               `json:"name,omitempty" bson:"name,omitempty" validate:"required,max=16"`
	Oid       int                  `json:"id,omitempty" bson:"id,omitempty"`
	Inbox     bool                 `json:"inbox" bson:"inbox"`
	TaskIDs   []primitive.ObjectID `json:"taskIDs,omitempty" bson:"taskIDs,omitempty"`
	Completed bool                 `json:"completed" bson:"completed"`
}

type ActionableUpdate struct {
	Name      string `json:"name,omitempty" validate:"required,max=20" bson:"name,omitempty"`
	Oid       *int   `json:"id,omitempty" bson:"id,omitempty"`
	Inbox     bool   `json:"inbox" bson:"inbox"`
	Completed *bool  `json:"completed,omitempty" bson:"completed,omitempty"`
}
