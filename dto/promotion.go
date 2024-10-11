package dto

import (
	"time"
)

type CreatePromotionRequest struct {
	Name          string    `json:"name" valid:"required"`
	DiscountType  string    `json:"discount_type" valid:"required"`
	DiscountValue float64   `json:"discount_value" valid:"required"`
	StartDate     time.Time `json:"start_date" valid:"required"`
	EndDate       time.Time `json:"end_date" valid:"required"`
	ApplicableTo  []string  `json:"applicable_to" valid:"required"`
}

type GetPromotion struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	DiscountType  string    `json:"discount_type"`
	DiscountValue float64   `json:"discount_value"`
	ApplicableTo  []string  `json:"applicable_to"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	IsActive      bool      `json:"is_active"`
}

type GetAllPromotionResponse struct {
	StatusCode int            `json:"statusCode"`
	Message    string         `json:"message"`
	Data       []GetPromotion `json:"data"`
}

type GetOnePromotionResponse struct {
	StatusCode int          `json:"statusCode"`
	Message    string       `json:"message"`
	Data       GetPromotion `json:"data"`
}

type ApplyPromotionRequest struct {
	ProductID   string `json:"product_id" valid:"required"`
	PromotionID string `json:"promotion_id" valid:"required"`
}

type ApplyPromotionResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
