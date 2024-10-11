package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/review_repository"
	"net/http"
)

type reviewService struct {
	reviewRepo review_repository.Repository
}

type ReviewService interface {
	ReviewCreate(reviewPayload dto.NewReviewRequest) (*dto.GetOneReviewsResponse, errs.MessageErr)
	GetReview() (*dto.GetReviewsResponse, errs.MessageErr)
}

func NewReviewService(reviewRepo review_repository.Repository) ReviewService {
	return &reviewService{
		reviewRepo: reviewRepo,
	}
}

func (rs *reviewService) ReviewCreate(reviewPayload dto.NewReviewRequest) (*dto.GetOneReviewsResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(reviewPayload)

	if err != nil {
		return nil, err
	}

	review := entity.Review{
		ProductID: reviewPayload.ProductId,
		Rating:    reviewPayload.Rating,
		Comment:   reviewPayload.Comment,
	}

	result, err := rs.reviewRepo.CreateReview(review)

	if err != nil {
		return nil, err
	}

	response := dto.GetOneReviewsResponse{
		StatusCode: http.StatusCreated,
		Message:    "Succesfully Create Reviews",
		Reviews: dto.GetReview{
			Id:        result.Id,
			ProductId: result.ProductID,
			Rating:    result.Rating,
			Comment:   result.Comment,
			CreatedAt: result.CreatedAt,
		},
	}

	return &response, nil

}

func (rs *reviewService) GetReview() (*dto.GetReviewsResponse, errs.MessageErr) {

	reviewResult, err := rs.reviewRepo.GetAllReview()

	if err != nil {
		return nil, err
	}

	reviews := []dto.GetReview{}

	for _, eachReview := range *reviewResult {
		review := dto.GetReview{
			Id:        eachReview.Id,
			ProductId: eachReview.ProductID,
			Rating:    eachReview.Rating,
			Comment:   eachReview.Comment,
			CreatedAt: eachReview.CreatedAt,
		}
		reviews = append(reviews, review)
	}

	response := dto.GetReviewsResponse{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Read Reviews Data",
		Reviews:    reviews,
	}
	return &response, nil
}
