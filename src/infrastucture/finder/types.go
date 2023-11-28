package finder

import (
	"github.com/flexuxs/clubHubApp/src/api/config"
	client_mongo "github.com/flexuxs/clubHubApp/src/api/fxmodule/client"
	interfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
)

type CompanyFinder struct {
	MongoClient *client_mongo.HotelMongoClient
	Config      *config.Configuration
}

func NewCompanyFinder(mongoClient *client_mongo.HotelMongoClient, config *config.Configuration) interfaces.ICompanyFinder {
	return &CompanyFinder{
		MongoClient: mongoClient,
		Config:      config,
	}
}
