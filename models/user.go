package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" validate:"required,min=2,max=10" bson:"name"`
	Password string             `json:"password" validate:"required" bson:"password"`
	ImgUrl   string             `json:"img_url" validate:"required" bson:"img_url"`
}
