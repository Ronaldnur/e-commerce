package entity

import "time"

type Product struct {
	Id           string    `bson:"_id"`
	Name         string    `bson:"name"`
	Price        float64   `bson:"price"`
	Stock        int       `bson:"stock"`
	UserId       string    `bson:"seller_id"`
	Promotion_Id *string   `bson:"promotion_id"`
	Created_at   time.Time `bson:"created_at"`
	Updated_at   time.Time `bson:"updated_at"`
}
