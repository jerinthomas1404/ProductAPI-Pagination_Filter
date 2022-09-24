package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitemtpty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Image       string             `json:"image,omitempty" bson:"image,omitemtpty"`
	Price       int                `json:"price,omitempty" bson:"price,omitemtpty"`
}
