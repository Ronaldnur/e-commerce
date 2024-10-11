package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentRequest struct {
	Amount     float64 `json:"amount" valid:"required"`
	Commission float64 `json:"commission" valid:"required"`
	Tax        float64 `json:"tax" valid:"required"`
}

type GetPayment struct {
	Id         primitive.ObjectID `json:"id"`
	User_Id    string             `json:"seller_id"`
	Amount     float64            `json:"amount"`
	Commission float64            `json:"commission"`
	Tax        float64            `json:"tax"`
	Status     string             `json:"status"`
}

type GetPaymentResponse struct {
	StatusCode int        `json:"statusCode"`
	Message    string     `json:"message"`
	Data       GetPayment `json:"data"`
}
type BalanceResponse struct {
	Id        primitive.ObjectID `json:"id"`
	User_Id   string             `json:"seller_id"`
	Total     float64            `json:"total"`
	Available float64            `json:"available"`
	Withdrawn float64            `json:"withdrawn"`
}

type GetBalanceResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       BalanceResponse `json:"data"`
}

type GetAllPaymentResponse struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       []GetPayment `json:"data"`
}
