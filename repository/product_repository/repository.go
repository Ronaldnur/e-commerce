package product_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreateProduct(productPayload entity.Product, userId string) errs.MessageErr
	FindAllProduct() ([]entity.Product, errs.MessageErr)
	FindProductById(productId string) (*entity.Product, errs.MessageErr)
	UpdateProduct(productId string, payload entity.Product) errs.MessageErr
	DeleteProduct(productId string) errs.MessageErr
	UpdateProductStock(productID string, newStock int) errs.MessageErr
}
