package review_mg

import (
	"context"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/review_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type reviewMg struct {
	db *mongo.Client
}

func NewReviewMg(db *mongo.Client) review_repository.Repository {
	return &reviewMg{
		db: db,
	}
}

func (rm *reviewMg) CreateReview(payload entity.Review) (*entity.Review, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := rm.db.Database("tokosehat").Collection("reviews")

	currentTime := time.Now()
	payload.CreatedAt = currentTime
	result, err := collection.InsertOne(ctx, payload)

	if err != nil {
		return nil, errs.NewInternalServerError("failed to insert reviews data")
	}

	payload.Id = result.InsertedID.(primitive.ObjectID)

	return &payload, nil
}

func (rm *reviewMg) GetAllReview() (*[]entity.Review, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := rm.db.Database("tokosehat").Collection("reviews")

	var reviews []entity.Review

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		// Mengembalikan error jika query gagal
		return nil, errs.NewInternalServerError("Error while fetching reviews")
	}

	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, errs.NewInternalServerError("Error while decoding reviews")
	}

	return &reviews, nil
}
