package user_mg

import (
	"context"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/user_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userMg struct {
	db *mongo.Client
}

func NewUserMg(db *mongo.Client) user_repository.Repository {
	return &userMg{
		db: db,
	}
}

func (u *userMg) GetUserByEmail(userEmail string) (*entity.User, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := u.db.Database("tokosehat").Collection("users")

	filter := bson.M{"email": userEmail}

	var user entity.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewNotFoundError("User Email Not Found")
		}
		return nil, errs.NewInternalServerError("Failed Get data from database")
	}
	return &user, nil
}

func (u *userMg) CreateUser(userPayload entity.User) (string, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := u.db.Database("tokosehat").Collection("users")

	var userId string

	user := bson.M{
		"email":      userPayload.Email,
		"username":   userPayload.Username,
		"password":   userPayload.Password,
		"verified":   userPayload.Verified,
		"address":    userPayload.Address,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	insertResult, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", errs.NewInternalServerError("Internal Server Error")
	}
	userId = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return userId, nil
}

func (u *userMg) GetUserByUsername(userUsername string) (*entity.User, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := u.db.Database("tokosehat").Collection("users")

	filter := bson.M{"username": userUsername}

	var user entity.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.NewNotFoundError("User Email Not Found")
		}
		return nil, errs.NewInternalServerError("Failed Get data from database")
	}
	return &user, nil
}
