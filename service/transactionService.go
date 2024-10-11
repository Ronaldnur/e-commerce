package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/transaction_repository"
	"net/http"
)

type transactionService struct {
	TransactionRepo transaction_repository.Repository
}

type TransactionService interface {
	CreatePayment(userId string, paymentPayload dto.PaymentRequest) (*dto.GetPaymentResponse, errs.MessageErr)
	GetSellerByPayments(sellerID string) (*dto.GetAllPaymentResponse, errs.MessageErr)
	GetSellerByBalance(sellerID string) (*dto.GetBalanceResponse, errs.MessageErr)
	WithdrawBalance(sellerID string, amount float64) (*dto.GetBalanceResponse, errs.MessageErr)
	CreateBalance(userId string) (*dto.GetBalanceResponse, errs.MessageErr)
}

func NewTransactionService(TransactionRepo transaction_repository.Repository) TransactionService {
	return &transactionService{
		TransactionRepo: TransactionRepo,
	}
}

func (ts *transactionService) CreatePayment(userId string, paymentPayload dto.PaymentRequest) (*dto.GetPaymentResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(paymentPayload)

	if err != nil {
		return nil, err
	}

	payment := entity.Payment{
		Amount:     paymentPayload.Amount,
		Commission: paymentPayload.Commission,
		Tax:        paymentPayload.Tax,
		Status:     "Pending",
	}

	makepayment, err := ts.TransactionRepo.CreatePayment(userId, &payment)

	if err != nil {
		return nil, err
	}

	response := dto.GetPaymentResponse{
		StatusCode: http.StatusCreated,
		Message:    "Succesfully Create Payment",
		Data: dto.GetPayment{
			Id:         makepayment.Id,
			User_Id:    makepayment.User_Id,
			Amount:     makepayment.Amount,
			Commission: makepayment.Commission,
			Tax:        makepayment.Tax,
			Status:     makepayment.Status,
		},
	}

	return &response, nil

}
func (ts *transactionService) GetSellerByPayments(sellerID string) (*dto.GetAllPaymentResponse, errs.MessageErr) {
	payments, err := ts.TransactionRepo.GetSellerPayments(sellerID)

	if err != nil {
		return nil, err
	}

	var paymentResponses []dto.GetPayment

	for _, payment := range *payments {
		paymentResponses = append(paymentResponses, dto.GetPayment{
			Id:         payment.Id,
			User_Id:    payment.User_Id,
			Amount:     payment.Amount,
			Commission: payment.Commission,
			Tax:        payment.Tax,
			Status:     payment.Status,
		})
	}

	response := dto.GetAllPaymentResponse{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Get Payments Data",
		Data:       paymentResponses,
	}

	return &response, nil
}

func (ts *transactionService) GetSellerByBalance(sellerID string) (*dto.GetBalanceResponse, errs.MessageErr) {
	balance, err := ts.TransactionRepo.GetSellerBalance(sellerID)

	if err != nil {
		return nil, err
	}

	response := dto.GetBalanceResponse{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Get Balance Data",
		Data: dto.BalanceResponse{
			Id:        balance.Id,
			User_Id:   balance.User_Id,
			Total:     balance.Total,
			Available: balance.Available,
			Withdrawn: balance.Withdrawn,
		},
	}
	return &response, nil

}

func (ts *transactionService) WithdrawBalance(sellerID string, amount float64) (*dto.GetBalanceResponse, errs.MessageErr) {
	balance, err := ts.TransactionRepo.GetSellerBalance(sellerID)

	if err != nil {
		return nil, err
	}

	if amount > balance.Available {
		return nil, errs.NewBadRequest("insufficient balance")
	}

	balance.Available -= amount
	balance.Withdrawn += amount

	err = ts.TransactionRepo.UpdateBalance(balance)

	if err != nil {
		return nil, err
	}

	response := dto.GetBalanceResponse{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Withdraw",
		Data: dto.BalanceResponse{
			Id:        balance.Id,
			User_Id:   balance.User_Id,
			Total:     balance.Total,
			Available: balance.Available,
			Withdrawn: balance.Withdrawn,
		},
	}

	return &response, nil
}

func (ts *transactionService) CreateBalance(userId string) (*dto.GetBalanceResponse, errs.MessageErr) {
	balance := entity.Balance{
		User_Id:   userId,
		Total:     1000.0,
		Available: 2000.0,
		Withdrawn: 0,
	}

	result, err := ts.TransactionRepo.CreateBalance(userId, &balance)

	if err != nil {
		return nil, err
	}

	response := dto.GetBalanceResponse{
		StatusCode: http.StatusOK,
		Message:    "Succesfully Withdraw",
		Data: dto.BalanceResponse{
			Id:        result.Id,
			User_Id:   result.User_Id,
			Total:     result.Total,
			Available: result.Available,
			Withdrawn: result.Withdrawn,
		},
	}

	return &response, nil
}
