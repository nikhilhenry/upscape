package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tag struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name,omitempty" validate:"required,max=10,alpha" bson:"name,omitempty"`
	Color string             `json:"color,omitempty" validate:"required,hexcolor" bson:"color,omitempty"`
}
