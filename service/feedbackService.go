package service

import (
	"mongo-api/dto"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/pkg/helpers"
	"mongo-api/repository/feedback_repository"
	"net/http"
)

type feedbackService struct {
	feedbackRepo feedback_repository.Repository
}

type FeedbackService interface {
	FeedbackCreate(userId string, createPayloadFeedback dto.NewFeedbackRequest) (*dto.GetFeedbackResponse, errs.MessageErr)
	GetAllFeedbackTicketData(userId string) (*dto.GetAllFeedbackResponse, errs.MessageErr)
}

func NewFeedbackService(feedbackRepo feedback_repository.Repository) FeedbackService {
	return &feedbackService{
		feedbackRepo: feedbackRepo,
	}
}

func (fs *feedbackService) FeedbackCreate(userId string, createPayloadFeedback dto.NewFeedbackRequest) (*dto.GetFeedbackResponse, errs.MessageErr) {
	err := helpers.ValidateStruct(createPayloadFeedback)

	if err != nil {
		return nil, err
	}

	feedback := entity.Feedback{
		Subject:     createPayloadFeedback.Subject,
		Description: createPayloadFeedback.Description,
	}

	createFeedback, err := fs.feedbackRepo.CreateFeedback(userId, feedback)

	if err != nil {
		return nil, err
	}

	response := dto.GetFeedbackResponse{
		StatusCode: http.StatusCreated,
		Message:    "Feedback successfully created",
		Data: dto.GetFeedback{
			Id:          createFeedback.Id,
			User_Id:     createFeedback.User_Id,
			Subject:     createFeedback.Subject,
			Description: createFeedback.Description,
			Created_At:  createFeedback.Created_At,
			Updated_At:  createFeedback.Updated_At,
		},
	}

	return &response, nil
}

func (fs *feedbackService) GetAllFeedbackTicketData(userId string) (*dto.GetAllFeedbackResponse, errs.MessageErr) {
	feedbacks, err := fs.feedbackRepo.GetAllFeedbackTicketData(userId)
	if err != nil {
		return nil, err
	}

	var feedbackResults []dto.GetFeedbackWithTicket

	for _, eachFeedback := range *feedbacks {
		feedback := dto.GetFeedbackWithTicket{
			Id:          eachFeedback.Feedback.Id,
			User_Id:     eachFeedback.Feedback.User_Id,
			Subject:     eachFeedback.Feedback.Subject,
			Description: eachFeedback.Feedback.Description,
			Created_At:  eachFeedback.Ticket.Created_At,
			Updated_At:  eachFeedback.Feedback.Updated_At,
			Tickets: dto.GetTicket{
				Id:         eachFeedback.Ticket.Id,
				Status:     eachFeedback.Ticket.Status,
				Created_At: eachFeedback.Ticket.Created_At,
				Updated_At: eachFeedback.Ticket.Updated_At,
			},
		}

		feedbackResults = append(feedbackResults, feedback)
	}

	response := dto.GetAllFeedbackResponse{

		StatusCode: http.StatusOK,
		Message:    "Successfuly Read Feedback With Ticket Data",
		Data:       feedbackResults,
	}
	return &response, nil
}
