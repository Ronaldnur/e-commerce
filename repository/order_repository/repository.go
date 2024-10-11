package order_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreateOrder(payload []entity.Order) errs.MessageErr
	GetSellerOrder() ([]entity.Order, errs.MessageErr)
	UpdateStatus(orderId string, status string) errs.MessageErr
}
