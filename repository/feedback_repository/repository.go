package feedback_repository

import (
	"mongo-api/entity"
	"mongo-api/pkg/errs"
)

type Repository interface {
	CreateFeedback(userId string, payloadFeedback entity.Feedback) (*entity.Feedback, errs.MessageErr)
	GetAllFeedbackTicketData(userId string) (*[]FeedbackWithTicket, errs.MessageErr)
}
