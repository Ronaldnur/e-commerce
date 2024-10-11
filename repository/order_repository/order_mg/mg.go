package order_mg

import (
	"context"
	"fmt"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/order_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderMg struct {
	db *mongo.Client
}

func NewOrderMg(db *mongo.Client) order_repository.Repository {
	return &orderMg{
		db: db,
	}
}

func (o *orderMg) CreateOrder(payload []entity.Order) errs.MessageErr {
	// Context dengan timeout untuk operasi MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := o.db.Database("tokosehat").Collection("orders")

	var orders []interface{}
	for _, order := range payload {
		orderDoc := bson.M{
			"product_id": order.Product_Id,
			"quantity":   order.Quantity,
			"status":     order.Status,
			"created_at": time.Now(),
			"updated_at": time.Now(),
		}
		orders = append(orders, orderDoc)
	}

	// Insert produk ke MongoDB
	_, err := collection.InsertMany(ctx, orders)
	if err != nil {
		fmt.Println(err)
		return errs.NewInternalServerError("failed to create orders in database")
	}
	return nil
}

func (o *orderMg) GetSellerOrder() ([]entity.Order, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := o.db.Database("tokosehat").Collection("orders")

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, errs.NewInternalServerError("Error fetching orders: ")
	}
	defer cursor.Close(ctx)

	var orders []entity.Order
	for cursor.Next(ctx) {
		var order entity.Order
		if err := cursor.Decode(&order); err != nil {
			return nil, errs.NewInternalServerError("Error decoding order: ")
		}
		orders = append(orders, order)
	}

	if err := cursor.Err(); err != nil {
		return nil, errs.NewInternalServerError("Cursor error: ")
	}
	return orders, nil
}

func (o *orderMg) UpdateStatus(orderId string, status string) errs.MessageErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		return errs.NewNotFoundError("invalid product ID")
	}

	collection := o.db.Database("tokosehat").Collection("orders")

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{"status": status},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {

		return errs.NewInternalServerError("Failed to update order status")
	}
	return nil
}
