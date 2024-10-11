package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID primitive.ObjectID `bson:"product_id,omitempty"`
	Rating    int                `bson:"rating,omitempty"`
	Comment   string             `bson:"comment,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}
