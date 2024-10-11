package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Feedback_Id primitive.ObjectID `bson:"feedback_id,omitempty"`
	User_Id     string             `bson:"seller_id"`
	Status      string             `bson:"status"`
	Created_At  time.Time          `bson:"created_at"`
	Updated_At  time.Time          `bson:"updated_at"`
}
