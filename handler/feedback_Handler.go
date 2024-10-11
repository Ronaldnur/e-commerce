package handler

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type feedbackHandler struct {
	FeedbackService service.FeedbackService
}

func NewFeedbackHandler(FeedbackService service.FeedbackService) feedbackHandler {
	return feedbackHandler{
		FeedbackService: FeedbackService,
	}
}

func (fh *feedbackHandler) FeedbackSupport(ctx *gin.Context) {
	var newFeedbackRequest dto.NewFeedbackRequest

	if err := ctx.ShouldBindJSON(&newFeedbackRequest); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)
	result, err := fh.FeedbackService.FeedbackCreate(user.Id, newFeedbackRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (fh *feedbackHandler) GetFeedback(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := fh.FeedbackService.GetAllFeedbackTicketData(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (fh *feedbackHandler) GetECHO(ctx echo.Context) {

}
