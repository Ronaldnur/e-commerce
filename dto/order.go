package dto

import "time"

type OrderItem struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type NewOrderRequest struct {
	Items  []OrderItem `json:"orders" valid:"required"`
	Status string      `json:"status"`
}

type NewOrderResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type GetAllOrder struct {
	Id         string    `json:"id"`
	Product_Id string    `json:"product_id"`
	User_Id    string    `json:"user_id"`
	Quantity   int       `json:"quantity"`
	Status     string    `json:"status"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type GetOrderResponse struct {
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       []GetAllOrder `json:"data"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" valid:"required"`
}

type UpdateStatusResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
