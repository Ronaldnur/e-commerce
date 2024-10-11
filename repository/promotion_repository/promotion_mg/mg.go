package promotion_mg

import (
	"context"
	"fmt"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/promotion_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type promotionMg struct {
	db *mongo.Client
}

func NewPromotionMg(db *mongo.Client) promotion_repository.Repository {
	return &promotionMg{
		db: db,
	}
}

func (p *promotionMg) CreatePromotion(payload entity.Promotion) (*entity.Promotion, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := p.db.Database("tokosehat").Collection("promotion")

	newPromotion := bson.M{
		"name":           payload.Name,
		"discount_type":  payload.DiscountType,
		"discount_value": payload.DiscountValue,
		"start_date":     payload.StartDate,
		"end_date":       payload.EndDate,
		"applicable_to":  payload.ApplicableTo,
		"is_active":      true,
	}

	result, err := collection.InsertOne(ctx, newPromotion)

	if err != nil {
		fmt.Println(err)
		return nil, errs.NewInternalServerError("failed to insert analytics data")
	}

	payload.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return &payload, nil
}

func (p *promotionMg) GetPromotionData() (*[]entity.Promotion, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := p.db.Database("tokosehat").Collection("promotion")
	// Define a variable to hold the result

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errs.NewInternalServerError("Error Fetching Promotion")
	}
	defer cursor.Close(ctx)

	var promotions []entity.Promotion

	// Iterate through the cursor and decode each document into the promotions slice
	for cursor.Next(ctx) {
		var promotion entity.Promotion
		if err := cursor.Decode(&promotion); err != nil {
			return nil, errs.NewInternalServerError("Error decoding Promotions")
		}
		promotions = append(promotions, promotion)
	}

	if err := cursor.Err(); err != nil {
		return nil, errs.NewInternalServerError("error iterating promotion")
	}

	return &promotions, nil
}

func (p *promotionMg) GetPromotionById(promotionId string) (*entity.Promotion, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(promotionId)
	if err != nil {
		return nil, errs.NewNotFoundError("invalid product ID")
	}
	collection := p.db.Database("tokosehat").Collection("promotion")

	// Define a variable to hold the result
	var promotion entity.Promotion

	// Find the promotion document, you can modify the filter as needed
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&promotion)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
			return nil, errs.NewNotFoundError("Promotions not found")
		}
		return nil, errs.NewInternalServerError("Error fetching promotions")
	}

	return &promotion, nil
}
