package client_mongo

import (
	"context"
	"log"
	"time"

	"github.com/flexuxs/clubHubApp/src/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelMongoClient struct {
	Client *mongo.Client
}

func NewHotelMongoClient(config *config.Configuration) *HotelMongoClient {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Replace "mydatabase" with your actual database name
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:admin@localhost:27017/hoteldb"))
	if err != nil {
		log.Print("Error connecting to MongoDB!!!!!!!!!!!")
		log.Fatal(err)
	}
	return &HotelMongoClient{Client: client}
}
