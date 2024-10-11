package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/order_repository"
	"mongo-api/repository/product_repository"
	"net/http"
)

type orderService struct {
	orderRepo   order_repository.Repository
	productRepo product_repository.Repository
}
type OrderService interface {
	CreateOrder(orderPayload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllSellerOrder() (*dto.GetOrderResponse, errs.MessageErr)
	UpdateStatus(orderId string, statusPayload dto.UpdateStatusRequest) (*dto.UpdateStatusResponse, errs.MessageErr)
}

func NewOrderService(orderRepo order_repository.Repository, productRepo product_repository.Repository) OrderService {
	return &orderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (os *orderService) CreateOrder(orderPayload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(orderPayload)

	if err != nil {
		return nil, err
	}

	createOrders := []entity.Order{}

	for _, eachOrder := range orderPayload.Items {

		product, err := os.productRepo.FindProductById(eachOrder.ProductID)

		if err != nil {
			if err.Status() == http.StatusNotFound {
				return nil, errs.NewNotFoundError("Product not Found")
			}
			return nil, err
		}

		if eachOrder.Quantity > product.Stock {
			return nil, errs.NewUnprocessibleEntityError("Insufficient stock for the requested quantity")
		}

		order := entity.Order{
			Product_Id: eachOrder.ProductID,
			Quantity:   eachOrder.Quantity,
			Status:     "pending",
		}

		createOrders = append(createOrders, order)

		product.Stock -= eachOrder.Quantity
		err = os.productRepo.UpdateProductStock(product.Id, product.Stock)

		if err != nil {
			return nil, err
		}
	}

	err = os.orderRepo.CreateOrder(createOrders)

	if err != nil {
		return nil, err
	}
	response := dto.NewOrderResponse{
		StatusCode: 201,
		Message:    "Order created successfully",
		Data:       createOrders,
	}

	return &response, nil
}

func (os *orderService) GetAllSellerOrder() (*dto.GetOrderResponse, errs.MessageErr) {

	orders, err := os.orderRepo.GetSellerOrder()

	if err != nil {
		return nil, err
	}

	var filteredOrders []dto.GetAllOrder

	for _, eachorder := range orders {
		order := dto.GetAllOrder{
			Id:         eachorder.Id,
			Product_Id: eachorder.Product_Id,
			Quantity:   eachorder.Quantity,
			User_Id:    eachorder.User_Id,
			Status:     eachorder.Status,
			Created_at: eachorder.Created_at,
			Updated_at: eachorder.Updated_at,
		}
		filteredOrders = append(filteredOrders, order)
	}

	response := dto.GetOrderResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Read Orders",
		Data:       filteredOrders,
	}

	return &response, nil
}

func (os *orderService) UpdateStatus(orderId string, statusPayload dto.UpdateStatusRequest) (*dto.UpdateStatusResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(statusPayload)

	if err != nil {
		return nil, err
	}
	if statusPayload.Status != "shipping" && statusPayload.Status != "canceled" {
		return nil, errs.NewBadRequest("Invalid status, It should be either 'shipping' or 'canceled'")
	}

	updateStatus := statusPayload.Status

	err = os.orderRepo.UpdateStatus(orderId, updateStatus)
	if err != nil {
		return nil, err
	}

	response := dto.UpdateStatusResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Update Product",
	}
	return &response, nil
}
