package review_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreateReview(payload entity.Review) (*entity.Review, errs.MessageErr)
	GetAllReview() (*[]entity.Review, errs.MessageErr)
}
