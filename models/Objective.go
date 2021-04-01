package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Objective struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name         string             `json:"name,omitempty" validate:"required,max=20" bson:"name,omitempty"`
	Completed    bool               `json:"completed" bson:"completed"`
	Body         string             `json:"body,omitempty" validate:"max=100" bson:"body,omitempty"`
	ScheduledFor time.Time          `json:"scheduled_for,omitempty" bson:"scheduled_for,omitempty"`
	CreatedAt    time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
