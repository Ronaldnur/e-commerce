package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewAnalyticsRequest struct {
	UserId             string `json:"seller_id" valid:"required"`
	TotalRevenue       int    `json:"total_revenue" valid:"required"`
	TotalOrders        int    `json:"total_orders" valid:"required"`
	BestSellingProduct string `json:"best_selling_product" valid:"required"`
}

type GetAnalytics struct {
	Id                 primitive.ObjectID `json:"_id,omitempty"`
	UserId             string             `json:"seller_id"`
	TotalRevenue       int                `json:"total_revenue"`
	TotalOrders        int                `json:"total_orders"`
	BestSellingProduct string             `json:"best_selling_product"`
	CreatedAt          time.Time          `json:"created_at"`
	Updated_at         time.Time          `json:"updated_at"`
}

type GetAnalyticsResponse struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       GetAnalytics `json:"Data"`
}

type GetAllAnalyticsResponse struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       []GetAnalytics `json:"Data"`
}
