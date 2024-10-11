package dto

import "time"

type NewProductRequest struct {
	Name  string  `json:"name" valid:"required"`
	Price float64 `json:"price" valid:"required"`
	Stock int     `json:"stock" valid:"required"`
}

type NewProductResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type GetAllProduct struct {
	Id           string    `json:"id"`
	Name         string    `json:"name" valid:"required"`
	Price        float64   `json:"price" valid:"required"`
	Stock        int       `json:"stock" valid:"required"`
	Promotion_Id *string   `bson:"promotion_id"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`
}

type GetProductResponse struct {
	Result     string          `json:"result"`
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       []GetAllProduct `json:"data_product"`
}
type GetOneProductResponse struct {
	Result     string        `json:"result"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Data       GetAllProduct `json:"data_product"`
}

type DeleteProductResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock" valid:"required"`
}

type UpdateStockResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
