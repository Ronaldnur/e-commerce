package transaction_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreatePayment(userId string, payment *entity.Payment) (*entity.Payment, errs.MessageErr)
	GetSellerPayments(userId string) (*[]entity.Payment, errs.MessageErr)
	GetSellerBalance(userId string) (*entity.Balance, errs.MessageErr)
	CreateBalance(userId string, balancePayload *entity.Balance) (*entity.Balance, errs.MessageErr)
	UpdateBalance(balance *entity.Balance) errs.MessageErr
}
