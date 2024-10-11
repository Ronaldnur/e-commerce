package feedback_mg

import (
	"context"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/feedback_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type feedbackMg struct {
	db *mongo.Client
}

func NewFeedbackMg(db *mongo.Client) feedback_repository.Repository {
	return &feedbackMg{
		db: db,
	}
}

func (f *feedbackMg) CreateFeedback(userId string, payloadFeedback entity.Feedback) (*entity.Feedback, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := f.db.Database("tokosehat").Collection("feedback")

	currentTime := time.Now()

	payloadFeedback.Created_At = currentTime
	payloadFeedback.Updated_At = currentTime
	payloadFeedback.User_Id = userId

	result, err := collection.InsertOne(ctx, payloadFeedback)
	if err != nil {
		return nil, errs.NewInternalServerError("failed to insert analytics data")
	}

	payloadFeedback.Id = result.InsertedID.(primitive.ObjectID)

	// Now create a ticket related to the feedback
	ticketCollection := f.db.Database("tokosehat").Collection("ticket")

	// Prepare the ticket payload
	ticket := entity.Ticket{
		Feedback_Id: payloadFeedback.Id, // Link the feedback ID to the ticket
		User_Id:     userId,
		Status:      "open",
		Created_At:  currentTime,
		Updated_At:  currentTime,
	}

	// Insert the ticket into the ticket collection
	_, err = ticketCollection.InsertOne(ctx, ticket)
	if err != nil {
		return nil, errs.NewInternalServerError("failed to create ticket")
	}

	return &payloadFeedback, nil
}

func (f *feedbackMg) GetAllFeedbackTicketData(userId string) (*[]feedback_repository.FeedbackWithTicket, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	feedbackCollection := f.db.Database("tokosehat").Collection("feedback")
	ticketCollection := f.db.Database("tokosehat").Collection("ticket")

	var feedbacks []entity.Feedback
	filter := bson.D{{Key: "seller_id", Value: userId}}
	// Fetch the cursor and handle the error
	cursor, err := feedbackCollection.Find(ctx, filter)
	if err != nil {
		return nil, errs.NewInternalServerError("error fetching feedback")
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &feedbacks); err != nil {
		return nil, errs.NewInternalServerError("error decoding feedback")
	}

	var feedbackWithTickets []feedback_repository.FeedbackWithTicket

	for _, feedback := range feedbacks {
		var ticket entity.Ticket // A single Ticket struct

		// Fetch the ticket for the current feedback
		ticketFilter := bson.D{{Key: "feedback_id", Value: feedback.Id}} // Assuming feedback.ID is the correct field
		ticketCursor, err := ticketCollection.Find(ctx, ticketFilter)
		if err != nil {
			return nil, errs.NewInternalServerError("error fetching tickets for feedback")
		}
		defer ticketCursor.Close(ctx) // Ensure the cursor is closed when done

		// Check if there is a ticket associated with the feedback
		if ticketCursor.Next(ctx) {
			if err := ticketCursor.Decode(&ticket); err != nil {
				return nil, errs.NewInternalServerError("error decoding ticket")
			}
		}

		// Combine feedback and the single ticket
		feedbackWithTickets = append(feedbackWithTickets, feedback_repository.FeedbackWithTicket{
			Feedback: feedback,
			Ticket:   ticket,
		})
	}

	return &feedbackWithTickets, nil
}
