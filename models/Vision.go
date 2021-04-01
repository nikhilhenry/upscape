package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vision struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" validate:"required,max=20,alphanum" bson:"name,omitempty"`
	ImgURL    string             `json:"img_url,omitempty" validate:"url" bson:"img_url"`
	Body      string             `json:"body,omitempty" validate:"required,max=100,alphanum" bson:"body,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
}
