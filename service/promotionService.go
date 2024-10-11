package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/product_repository"
	"mongo-api/repository/promotion_repository"
	"net/http"
)

type promotionService struct {
	promotionRepo promotion_repository.Repository
	productRepo   product_repository.Repository
}

type PromotionService interface {
	PromotionCreate(req dto.CreatePromotionRequest) (*dto.GetOnePromotionResponse, errs.MessageErr)
	ApplyPromotionProduct(applyPromotion dto.ApplyPromotionRequest) (*dto.ApplyPromotionResponse, errs.MessageErr)
	GetPromotionAllData() (*dto.GetAllPromotionResponse, errs.MessageErr)
}

func NewPromotionService(promotionRepo promotion_repository.Repository, productRepo product_repository.Repository) PromotionService {
	return &promotionService{
		promotionRepo: promotionRepo,
		productRepo:   productRepo,
	}
}

func (sp *promotionService) PromotionCreate(req dto.CreatePromotionRequest) (*dto.GetOnePromotionResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(req)

	if err != nil {
		return nil, err
	}

	promotion := entity.Promotion{
		Name:          req.Name,
		DiscountType:  req.DiscountType,
		DiscountValue: req.DiscountValue,
		StartDate:     req.StartDate,
		EndDate:       req.EndDate,
		ApplicableTo:  req.ApplicableTo,
		IsActive:      true, // Aktifkan promosi
	}

	promo, err := sp.promotionRepo.CreatePromotion(promotion)

	if err != nil {
		return nil, err
	}

	response := dto.GetOnePromotionResponse{
		StatusCode: http.StatusCreated,
		Message:    "Successfuly Create Promotions",
		Data: dto.GetPromotion{
			Id:            promo.Id,
			Name:          promo.Name,
			DiscountType:  promo.DiscountType,
			DiscountValue: promo.DiscountValue,
			ApplicableTo:  promo.ApplicableTo,
			StartDate:     promo.StartDate,
			EndDate:       promo.EndDate,
			IsActive:      true,
		},
	}

	return &response, nil
}

func (sp *promotionService) ApplyPromotionProduct(applyPromotion dto.ApplyPromotionRequest) (*dto.ApplyPromotionResponse, errs.MessageErr) {
	product, err := sp.productRepo.FindProductById(applyPromotion.ProductID)

	if err != nil {
		return nil, err
	}

	promotion, err := sp.promotionRepo.GetPromotionById(applyPromotion.PromotionID)
	if err != nil {
		return nil, err
	}

	if !isProductAvailableInPromotion(product.Id, promotion.ApplicableTo) {
		return nil, errs.NewNotFoundError("Product ID not available for this promotion")
	}
	finalPrice := product.Price

	// Apply discount based on the discount type
	if promotion.DiscountType == "percentage" {
		// Calculate percentage discount
		discountAmount := product.Price * (promotion.DiscountValue / 100)
		finalPrice -= discountAmount
	} else if promotion.DiscountType == "fixed" {
		// Apply fixed discount value
		finalPrice -= promotion.DiscountValue
	} else {
		return nil, errs.NewBadRequest("Invalid Discount Type")
	}

	if finalPrice < 0 {
		finalPrice = 0
	}

	product.Price = finalPrice
	product.Promotion_Id = &promotion.Id

	err = sp.productRepo.UpdateProduct(product.Id, *product)

	if err != nil {
		return nil, err
	}

	response := dto.ApplyPromotionResponse{
		StatusCode: http.StatusCreated,
		Message:    "Successfuly apply Promotions to product",
	}

	return &response, nil
}

func isProductAvailableInPromotion(productId string, availableProductIDs []string) bool {
	for _, id := range availableProductIDs {
		if id == productId {
			return true
		}
	}
	return false
}

func (sp *promotionService) GetPromotionAllData() (*dto.GetAllPromotionResponse, errs.MessageErr) {
	promotions, err := sp.promotionRepo.GetPromotionData()

	if err != nil {
		return nil, err
	}

	var promotionDTOs []dto.GetPromotion

	for _, promo := range *promotions {
		promotionDTOs = append(promotionDTOs, dto.GetPromotion{
			Id:            promo.Id,
			Name:          promo.Name,
			DiscountType:  promo.DiscountType,
			DiscountValue: promo.DiscountValue,
			ApplicableTo:  promo.ApplicableTo,
			StartDate:     promo.StartDate,
			EndDate:       promo.EndDate,
			IsActive:      promo.IsActive,
			// Add other fields as necessary
		})
	}

	response := dto.GetAllPromotionResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Create Promotions",
		Data:       promotionDTOs,
	}

	return &response, nil
}
