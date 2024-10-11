package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payment struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_Id    string             `bson:"seller_id" json:"seller_id"`
	Amount     float64            `bson:"amount" json:"amount"`
	Commission float64            `bson:"commission" json:"commission"`
	Tax        float64            `bson:"tax" json:"tax"`
	Status     string             `bson:"status" json:"status"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
}

type Balance struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	User_Id   string             `bson:"seller_id" json:"seller_id"`
	Total     float64            `bson:"total" json:"total"`
	Available float64            `bson:"available" json:"available"`
	Withdrawn float64            `bson:"withdrawn" json:"withdrawn"`
}
