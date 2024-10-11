package entity

import (
	"time"
)

type Promotion struct {
	Id            string    `bson:"_id"`
	Name          string    `bson:"name"`
	DiscountType  string    `bson:"discount_type"`
	DiscountValue float64   `bson:"discount_value"`
	ApplicableTo  []string  `bson:"applicable_to"`
	StartDate     time.Time `bson:"start_date"`
	EndDate       time.Time `bson:"end_date"`
	IsActive      bool      `bson:"is_active"`
}
