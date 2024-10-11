package handler

import (
	"mongo-api/dto"
	"mongo-api/pkg/errs"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

type reviewHandler struct {
	ReviewService service.ReviewService
}

func NewReviewHandler(ReviewService service.ReviewService) reviewHandler {
	return reviewHandler{
		ReviewService: ReviewService,
	}
}

func (rh *reviewHandler) CreateReview(ctx *gin.Context) {
	var newReview dto.NewReviewRequest

	if err := ctx.ShouldBindJSON(&newReview); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := rh.ReviewService.ReviewCreate(newReview)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (rh *reviewHandler) GetAllDataReview(ctx *gin.Context) {

	result, err := rh.ReviewService.GetReview()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
