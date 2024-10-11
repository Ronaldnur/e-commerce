package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewReviewRequest struct {
	ProductId primitive.ObjectID `json:"product_id" valid:"required"`
	Rating    int                `json:"rating" valid:"required"`
	Comment   string             `json:"comment" valid:"required"`
}

type GetReviewsResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Reviews    []GetReview `json:"reviews"`
}
type GetReview struct {
	Id        primitive.ObjectID `json:"review_id"`
	ProductId primitive.ObjectID `json:"product_id"`
	Rating    int                `json:"rating"`
	Comment   string             `json:"comment"`
	CreatedAt time.Time          `json:"created_at"`
}

type GetOneReviewsResponse struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Reviews    GetReview `json:"reviews"`
}
