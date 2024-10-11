package analyticsreport_mg

import (
	"context"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/analyticsreport_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type analytics_mg struct {
	db *mongo.Client
}

func NewAnalyticsMg(db *mongo.Client) analyticsreport_repository.Repository {
	return &analytics_mg{
		db: db,
	}
}

func (a *analytics_mg) CreateAnalytics(payloadReport entity.Analytics) (*entity.Analytics, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	currentTime := time.Now()
	payloadReport.Created_At = currentTime
	payloadReport.Updated_At = currentTime

	collection := a.db.Database("tokosehat").Collection("analytics")

	result, err := collection.InsertOne(ctx, payloadReport)

	if err != nil {
		return nil, errs.NewInternalServerError("failed to insert analytics data")
	}

	payloadReport.Id = result.InsertedID.(primitive.ObjectID)

	return &payloadReport, nil
}

func (a *analytics_mg) GetReport(userId string) ([]*entity.Analytics, errs.MessageErr) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := a.db.Database("tokosehat").Collection("analytics")

	filter := bson.M{"seller_id": userId}

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, errs.NewInternalServerError("Error when fetching analytics")
	}
	defer cursor.Close(ctx)

	var analyticsList []*entity.Analytics

	for cursor.Next(ctx) {
		var analytics entity.Analytics
		if err := cursor.Decode(&analytics); err != nil {
			return nil, errs.NewInternalServerError("Error when decoding analytics data")
		}
		analyticsList = append(analyticsList, &analytics)
	}

	if err := cursor.Err(); err != nil {
		return nil, errs.NewInternalServerError("Error during cursor iteration")
	}
	return analyticsList, nil
}
