package handler

import (
	"mongo-api/dto"
	"mongo-api/pkg/errs"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NeworderHandler(OrderService service.OrderService) orderHandler {
	return orderHandler{
		OrderService: OrderService,
	}
}

func (oh *orderHandler) MakeOrder(ctx *gin.Context) {

	var newOrder dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	result, err := oh.OrderService.CreateOrder(newOrder)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(result.StatusCode, result)
}

func (oh *orderHandler) GetOrder(ctx *gin.Context) {

	result, err := oh.OrderService.GetAllSellerOrder()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (oh *orderHandler) UpdateStatus(ctx *gin.Context) {
	orderId := ctx.Param("orderId")
	var updateStatus dto.UpdateStatusRequest

	if err := ctx.ShouldBindJSON(&updateStatus); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}
	result, err := oh.OrderService.UpdateStatus(orderId, updateStatus)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}
