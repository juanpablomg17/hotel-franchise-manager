package client_mongo

import (
	"context"
	"log"

	"github.com/flexuxs/clubHubApp/src/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type HotelMongoClient struct {
	Client *mongo.Client
}

func NewHotelMongoClient(config *config.Configuration) *HotelMongoClient {

	// mongodb://root:admin@localhost:27017/?directConnection=true&authMechanism=DEFAULT

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://root:admin@localhost:27017/hoteldb?ssl=false&authSource=admin"))
	if err != nil {
		log.Print("Error connecting to MongoDB!!!!!!!!!!!")
		log.Fatal(err)
	}
	return &HotelMongoClient{Client: client}
}
