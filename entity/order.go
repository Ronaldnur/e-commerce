package entity

import "time"

type Order struct {
	Id          string    `bson:"_id"`
	Product_Id  string    `bson:"product_id"`
	User_Id     string    `bson:"user_id"`
	Quantity    int       `bson:"quantity"`
	Status      string    `bson:"status"`
	Created_at  time.Time `bson:"created_at"`
	Updated_at  time.Time `bson:"updated_at"`
}
