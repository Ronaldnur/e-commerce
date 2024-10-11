package handler

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

type analyticsHandler struct {
	AnalyticsService service.AnalyticsService
}

func NewAnalyticsHandler(AnalyticsService service.AnalyticsService) analyticsHandler {
	return analyticsHandler{
		AnalyticsService: AnalyticsService,
	}
}

func (ah *analyticsHandler) CreateAnalytics(ctx *gin.Context) {

	var analyticsReport dto.NewAnalyticsRequest

	if err := ctx.ShouldBindJSON(&analyticsReport); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := ah.AnalyticsService.CreateAnalytics(analyticsReport)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (ah *analyticsHandler) GetSellerReport(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := ah.AnalyticsService.FindReportSellerId(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}
