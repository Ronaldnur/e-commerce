package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewFeedbackRequest struct {
	Subject     string `json:"subject" valid:"required"`
	Description string `json:"description" valid:"required"`
}

type GetFeedback struct {
	Id          primitive.ObjectID `json:"_id "`
	User_Id     string             `json:"seller_id"`
	Subject     string             `json:"subject"`
	Description string             `json:"description"`
	Created_At  time.Time          `json:"created_at"`
	Updated_At  time.Time          `json:"updated_at"`
}

type GetFeedbackResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       GetFeedback `json:"data"`
}

type GetAllFeedbackResponse struct {
	StatusCode int                     `json:"statusCode"`
	Message    string                  `json:"message"`
	Data       []GetFeedbackWithTicket `json:"data"`
}

type GetFeedbackWithTicket struct {
	Id          primitive.ObjectID `json:"_id "`
	User_Id     string             `json:"seller_id"`
	Subject     string             `json:"subject"`
	Description string             `json:"description"`
	Created_At  time.Time          `json:"created_at"`
	Updated_At  time.Time          `json:"updated_at"`
	Tickets     GetTicket          `json:"ticket"`
}

type GetTicket struct {
	Id         primitive.ObjectID `json:"id"`
	Status     string             `json:"status"`
	Created_At time.Time          `json:"created_at"`
	Updated_At time.Time          `json:"updated_at"`
}
