package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Objective struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required,max=20,alphanum" bson:"name,omitempty"`
	Completed    bool               `json:"completed,omitempty" validate:"required" bson:"completed,omitempty"`
	Body         string             `json:"body,omitempty" validate:"required,max=100,alphanum" bson:"body,omitempty"`
	ScheduledFor time.Time          `json:"scheduledFor,omitempty" bson:"scheduledFor,omitempty"`
	CreatedAt    time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
