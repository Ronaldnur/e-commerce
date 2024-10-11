package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/product_repository"
	"net/http"
)

type productService struct {
	productRepo product_repository.Repository
}

type ProductService interface {
	ProductCreate(payload *dto.NewProductRequest, userId string) (*dto.NewProductResponse, errs.MessageErr)
	GetProductData() (*dto.GetProductResponse, errs.MessageErr)
	GetOneProductData(productId string) (*dto.GetOneProductResponse, errs.MessageErr)
	UpdateProductById(productId string, updatePayload *dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	DeleteProduct(productId string) (*dto.DeleteProductResponse, errs.MessageErr)
	UpdateProductStock(productId string, stockPayload dto.UpdateStockRequest) (*dto.UpdateStockResponse, errs.MessageErr)
}

func NewProductService(productRepo product_repository.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (ps *productService) ProductCreate(payload *dto.NewProductRequest, userId string) (*dto.NewProductResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}

	product := entity.Product{
		Name:  payload.Name,
		Price: payload.Price,
		Stock: payload.Stock,
	}
	err = ps.productRepo.CreateProduct(product, userId)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Message:    "order successfully created",
		Data:       &product,
	}

	return &response, nil
}

func (ps *productService) GetProductData() (*dto.GetProductResponse, errs.MessageErr) {
	products, err := ps.productRepo.FindAllProduct()

	if err != nil {
		return nil, err
	}

	productResult := []dto.GetAllProduct{}

	for _, eachProduct := range products {
		product := dto.GetAllProduct{
			Id:           eachProduct.Id,
			Name:         eachProduct.Name,
			Price:        eachProduct.Price,
			Stock:        eachProduct.Stock,
			Promotion_Id: eachProduct.Promotion_Id,
			Created_at:   eachProduct.Created_at,
			Updated_at:   eachProduct.Updated_at,
		}
		productResult = append(productResult, product)
	}

	response := dto.GetProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Read Orders",
		Data:       productResult,
	}

	return &response, nil
}

func (ps *productService) GetOneProductData(productId string) (*dto.GetOneProductResponse, errs.MessageErr) {

	product, err := ps.productRepo.FindProductById(productId)
	if err != nil {
		return nil, err
	}

	response := dto.GetOneProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfully Read Product Data by ID",
		Data: dto.GetAllProduct{
			Id:           product.Id,
			Name:         product.Name,
			Price:        product.Price,
			Stock:        product.Stock,
			Promotion_Id: product.Promotion_Id,
			Created_at:   product.Created_at,
			Updated_at:   product.Updated_at,
		},
	}

	return &response, nil
}

func (ps *productService) UpdateProductById(productId string, updatePayload *dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(updatePayload)

	if err != nil {
		return nil, err
	}

	updateProduct := entity.Product{
		Name:  updatePayload.Name,
		Price: updatePayload.Price,
		Stock: updatePayload.Stock,
	}

	err = ps.productRepo.UpdateProduct(productId, updateProduct)
	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Update Product",
		Data:       nil,
	}

	return &response, nil

}

func (ps *productService) DeleteProduct(productId string) (*dto.DeleteProductResponse, errs.MessageErr) {

	err := ps.productRepo.DeleteProduct(productId)

	if err != nil {
		return nil, err
	}

	response := dto.DeleteProductResponse{
		StatusCode: http.StatusOK,
		Message:    "Your Product has been succsessfully deleted",
	}

	return &response, nil
}

func (ps *productService) UpdateProductStock(productId string, stockPayload dto.UpdateStockRequest) (*dto.UpdateStockResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(stockPayload)

	if err != nil {
		return nil, err
	}
	updateStock := stockPayload.Stock

	err = ps.productRepo.UpdateProductStock(productId, updateStock)

	if err != nil {
		return nil, err
	}

	response := dto.UpdateStockResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Update Stock Product",
		Data:       &updateStock,
	}

	return &response, nil
}
