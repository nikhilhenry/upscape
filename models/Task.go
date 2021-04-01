package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required,alphanum" bson:"name,omitempty"`
	Duration    int                `json:"duration,omitempty" bson:"duration,omitempty"`
	Completed   bool               `json:"completed,omitempty" validate:"required" bson:"completed"`
	TagIDs      []string           `json:"tag_ids" bson:"tag_ids"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	CompletedAt time.Time          `json:"completed_at" bson:"completed_at"`
}
