package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Client
)

const uri = "mongodb://localhost:27017/"

func connectDb() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the primary to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	fmt.Println("Successfully connected and pinged.")

	db = client
	return client

}

func InitiliazeDatabase() {
	db = connectDb()
}

func GetDatabaseInstance() *mongo.Client {
	return db
}

func GetCollection(collectionName string) *mongo.Collection {
	if db == nil {
		log.Fatal("Database not initialized")
	}
	collection := db.Database("tokosehat").Collection(collectionName)
	return collection
}
