package product_mg

import (
	"context"
	"fmt"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/product_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type productMg struct {
	db *mongo.Client
}

func NewProductMg(db *mongo.Client) product_repository.Repository {
	return &productMg{
		db: db,
	}
}

func (p *productMg) CreateProduct(productPayload entity.Product, userId string) errs.MessageErr {
	// Context dengan timeout untuk operasi MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := p.db.Database("tokosehat").Collection("product")

	// Konversi entity.Product ke BSON sebelum dimasukkan
	product := bson.M{
		"name":         productPayload.Name,
		"price":        productPayload.Price,
		"stock":        productPayload.Stock,
		"user_id":      userId,
		"promotion_id": productPayload.Promotion_Id,
		"created_at":   time.Now(),
		"updated_at":   time.Now(),
	}

	// Insert produk ke MongoDB
	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		return errs.NewInternalServerError("failed to create product in database")
	}

	return nil
}

func (p *productMg) FindAllProduct() ([]entity.Product, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := p.db.Database("tokosehat").Collection("product")

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, errs.NewInternalServerError("Failed to retrieve products from database")
	}

	defer cursor.Close(ctx)

	var products []entity.Product

	for cursor.Next(ctx) {
		var product entity.Product

		if err := cursor.Decode(&product); err != nil {
			fmt.Println(err)
			return nil, errs.NewInternalServerError("failed to decode product data")
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, errs.NewInternalServerError("error occured while iterating products")
	}
	return products, nil
}

func (p *productMg) FindProductById(productId string) (*entity.Product, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return nil, errs.NewNotFoundError("invalid product ID")
	}
	collection := p.db.Database("tokosehat").Collection("product")

	filter := bson.M{"_id": objectID}

	var product entity.Product

	err = collection.FindOne(ctx, filter).Decode(&product)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
			return nil, errs.NewInternalServerError("product not found")
		}
		return nil, errs.NewInternalServerError("Failed get data from database")
	}

	return &product, nil
}

func (p *productMg) UpdateProduct(productId string, payload entity.Product) errs.MessageErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return errs.NewNotFoundError("invalid product ID")
	}
	collection := p.db.Database("tokosehat").Collection("product")

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"name":       payload.Name,
			"price":      payload.Price,
			"stock":      payload.Stock,
			"updated_at": time.Now().UTC(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println(err)
			return errs.NewInternalServerError("product not found")
		}
		return errs.NewInternalServerError("Failed update data to database")
	}
	return nil
}

func (p *productMg) DeleteProduct(productId string) errs.MessageErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return errs.NewNotFoundError("invalid product ID")
	}
	collection := p.db.Database("tokosehat").Collection("product")

	filter := bson.M{"_id": objectID}

	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return errs.NewInternalServerError("failed to delete product")
	}

	// Check if any document was deleted
	if result.DeletedCount == 0 {
		return errs.NewNotFoundError("product not found")
	}
	return nil
}

func (p *productMg) UpdateProductStock(productId string, newStock int) errs.MessageErr {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return errs.NewNotFoundError("invalid product ID")
	}
	collection := p.db.Database("tokosehat").Collection("product")

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{"stock": newStock}, // Sesuaikan dengan field nama stok di model
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		// Tangani kesalahan jika update gagal
		return errs.NewInternalServerError("Failed to update product stock")
	}
	return nil
}
