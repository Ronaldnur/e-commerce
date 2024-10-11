package handler

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/service"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(TransactionService service.TransactionService) transactionHandler {
	return transactionHandler{
		TransactionService: TransactionService,
	}
}

func (th *transactionHandler) CreatePayment(ctx *gin.Context) {
	var newPayment dto.PaymentRequest

	if err := ctx.ShouldBindJSON(&newPayment); err != nil {
		errBindJson := errs.NewUnprocessibleEntityError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := ctx.MustGet("userData").(entity.User)

	result, err := th.TransactionService.CreatePayment(user.Id, newPayment)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *transactionHandler) GetPaymentSeller(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := th.TransactionService.GetSellerByPayments(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *transactionHandler) GetBalanceSeller(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := th.TransactionService.GetSellerByBalance(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *transactionHandler) Withdraw(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	amount, err := helpers.GetQueryFloat(ctx, "amount")

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	result, err := th.TransactionService.WithdrawBalance(user.Id, amount)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}
	ctx.JSON(result.StatusCode, result)
}

func (th *transactionHandler) CreateBalance(ctx *gin.Context) {
	user := ctx.MustGet("userData").(entity.User)

	result, err := th.TransactionService.CreateBalance(user.Id)
	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
	}

	ctx.JSON(result.StatusCode, result)
}
