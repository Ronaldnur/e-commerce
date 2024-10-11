package promotion_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreatePromotion(payload entity.Promotion) (*entity.Promotion, errs.MessageErr)
	GetPromotionData() (*[]entity.Promotion, errs.MessageErr)
	GetPromotionById(promotionId string) (*entity.Promotion, errs.MessageErr)
}
