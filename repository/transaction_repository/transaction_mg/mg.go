package transaction_mg

import (
	"context"
	"mongo-api/entity"
	"mongo-api/pkg/errs"
	"mongo-api/repository/transaction_repository"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type transactionMg struct {
	db *mongo.Client
}

func NewTransactionMg(db *mongo.Client) transaction_repository.Repository {
	return &transactionMg{
		db: db,
	}
}

func (t *transactionMg) CreatePayment(userId string, payment *entity.Payment) (*entity.Payment, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := t.db.Database("tokosehat").Collection("transaction")

	payment.User_Id = userId
	payment.CreatedAt = time.Now()
	payment.Status = "Pending"

	result, err := collection.InsertOne(ctx, payment)

	if err != nil {
		return nil, errs.NewInternalServerError("Gagal menyimpan pembayaran")
	}
	payment.Id = result.InsertedID.(primitive.ObjectID)
	return payment, nil
}

func (t *transactionMg) GetSellerPayments(userId string) (*[]entity.Payment, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := t.db.Database("tokosehat").Collection("transaction")

	filter := bson.M{"seller_id": userId}

	var payments []entity.Payment
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errs.NewInternalServerError("Gagal dalam mencari Seller")
	}

	// Iterate melalui cursor dan decode hasilnya ke slice of payments
	if err = cursor.All(ctx, &payments); err != nil {
		return nil, errs.NewInternalServerError("Gagal dalam mendecode pembayaran")
	}
	return &payments, nil
}

func (t *transactionMg) GetSellerBalance(userId string) (*entity.Balance, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := t.db.Database("tokosehat").Collection("balance")

	filter := bson.M{"seller_id": userId}
	var balance entity.Balance
	err := collection.FindOne(ctx, filter).Decode(&balance)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Jika saldo tidak ditemukan, inisialisasi dengan saldo 0
			balance = entity.Balance{
				Id:        balance.Id,
				User_Id:   userId,
				Available: 0,
				Withdrawn: 0,
			}
			return &balance, nil
		}
		return nil, errs.NewInternalServerError("Gagal dalam mencari balance seller ")
	}

	return &balance, nil
}

func (t *transactionMg) UpdateBalance(balance *entity.Balance) errs.MessageErr {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := t.db.Database("tokosehat").Collection("balance")

	filter := bson.M{"seller_id": balance.User_Id}
	update := bson.M{
		"$set": bson.M{
			"available": balance.Available,
			"withdrawn": balance.Withdrawn,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return errs.NewInternalServerError("Gagal dalam mengupdate balance ")
	}

	return nil
}
func (t *transactionMg) CreateBalance(userId string, balancePayload *entity.Balance) (*entity.Balance, errs.MessageErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Pilih database dan collection yang tepat
	collection := t.db.Database("tokosehat").Collection("balance")
	balancePayload.User_Id = userId

	result, err := collection.InsertOne(ctx, balancePayload)

	if err != nil {
		return nil, errs.NewInternalServerError("Gagal menyimpan pembayaran")
	}
	balancePayload.Id = result.InsertedID.(primitive.ObjectID)

	return balancePayload, nil
}
