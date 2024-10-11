package analyticsreport_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreateAnalytics(payloadReport entity.Analytics) (*entity.Analytics, errs.MessageErr)
	GetReport(userId string) ([]*entity.Analytics, errs.MessageErr)
}
