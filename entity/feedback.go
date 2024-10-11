package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Feedback struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	User_Id     string             `bson:"seller_id"`
	Subject     string             `bson:"subject"`
	Description string             `bson:"description"`
	Created_At  time.Time          `bson:"created_at"`
	Updated_At  time.Time          `bson:"updated_at"`
}
