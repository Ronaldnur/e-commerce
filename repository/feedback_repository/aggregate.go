package feedback_repository

import "mongo-api/entity"

type FeedbackWithTicket struct {
	Ticket   entity.Ticket
	Feedback entity.Feedback
}
