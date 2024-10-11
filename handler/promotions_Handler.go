package handler

import (
	"mongo-api/dto"
	"mongo-api/pkg/errs"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

type promotionHandler struct {
	PromotionService service.PromotionService
}

func NewPromotionHandler(PromotionService service.PromotionService) promotionHandler {
	return promotionHandler{
		PromotionService: PromotionService,
	}
}

func (p *promotionHandler) MakePromotion(ctx *gin.Context) {
	var newPromotion dto.CreatePromotionRequest

	if err := ctx.ShouldBindJSON(&newPromotion); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := p.PromotionService.PromotionCreate(newPromotion)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (p *promotionHandler) GetPromotion(ctx *gin.Context) {
	result, err := p.PromotionService.GetPromotionAllData()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (p *promotionHandler) ApplyPromotion(ctx *gin.Context) {
	var applyPromotion dto.ApplyPromotionRequest

	if err := ctx.ShouldBindJSON(&applyPromotion); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := p.PromotionService.ApplyPromotionProduct(applyPromotion)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
