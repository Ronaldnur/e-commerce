package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Analytics struct {
	Id                 primitive.ObjectID `bson:"_id,omitempty"`
	UserId             string             `bson:"seller_id"`
	TotalRevenue       int                `bson:"total_revenue"`
	TotalOrders        int                `bson:"total_orders"`
	BestSellingProduct string             `bson:"best_selling_product"`
	Created_At         time.Time          `bson:"created_at"`
	Updated_At         time.Time          `bson:"updated_at"`
}
