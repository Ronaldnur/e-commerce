package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/analyticsreport_repository"
	"net/http"
)

type analyticsService struct {
	analyticsRepo analyticsreport_repository.Repository
}

type AnalyticsService interface {
	CreateAnalytics(payload dto.NewAnalyticsRequest) (*dto.GetAnalyticsResponse, errs.MessageErr)
	FindReportSellerId(userId string) (*dto.GetAllAnalyticsResponse, errs.MessageErr)
}

func NewAnalyticsService(analyticsRepo analyticsreport_repository.Repository) AnalyticsService {
	return &analyticsService{
		analyticsRepo: analyticsRepo,
	}
}

func (as *analyticsService) CreateAnalytics(payload dto.NewAnalyticsRequest) (*dto.GetAnalyticsResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(payload)

	if err != nil {
		return nil, err
	}
	analytics := entity.Analytics{
		UserId:             payload.UserId,
		TotalRevenue:       payload.TotalRevenue,
		TotalOrders:        payload.TotalOrders,
		BestSellingProduct: payload.BestSellingProduct,
	}

	// Panggil fungsi repo untuk membuat data analytics
	createdAnalytics, err := as.analyticsRepo.CreateAnalytics(analytics)
	if err != nil {
		return nil, errs.NewInternalServerError("Failed to create analytics")
	}

	response := dto.GetAnalyticsResponse{
		StatusCode: http.StatusCreated,
		Message:    "Analyticts Report created successfully",
		Data: dto.GetAnalytics{
			Id:                 createdAnalytics.Id,
			UserId:             createdAnalytics.UserId,
			TotalRevenue:       createdAnalytics.TotalRevenue,
			TotalOrders:        createdAnalytics.TotalOrders,
			BestSellingProduct: createdAnalytics.BestSellingProduct,
			CreatedAt:          createdAnalytics.Created_At,
			Updated_at:         createdAnalytics.Updated_At,
		},
	}

	return &response, nil
}

func (as *analyticsService) FindReportSellerId(userId string) (*dto.GetAllAnalyticsResponse, errs.MessageErr) {

	getReport, err := as.analyticsRepo.GetReport(userId)

	if err != nil {
		return nil, err
	}

	var analyticsReports []dto.GetAnalytics

	for _, eachReport := range getReport {
		analyticsReport := dto.GetAnalytics{
			Id:                 eachReport.Id,
			UserId:             eachReport.UserId,
			TotalRevenue:       eachReport.TotalRevenue,
			TotalOrders:        eachReport.TotalOrders,
			BestSellingProduct: eachReport.BestSellingProduct,
			CreatedAt:          eachReport.Created_At,
			Updated_at:         eachReport.Updated_At,
		}
		analyticsReports = append(analyticsReports, analyticsReport)
	}

	response := dto.GetAllAnalyticsResponse{
		StatusCode: http.StatusOK,
		Message:    "Successfuly Read Analytics",
		Data:       analyticsReports,
	}

	return &response, nil
}
