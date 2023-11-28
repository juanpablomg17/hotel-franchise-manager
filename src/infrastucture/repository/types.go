package repository

import (
	"github.com/flexuxs/clubHubApp/src/api/config"
	client_mongo "github.com/flexuxs/clubHubApp/src/api/fxmodule/client"
	interfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
)

type CompanyRepository struct {
	MongoClient *client_mongo.HotelMongoClient
	Config      *config.Configuration
}

func NewCompanyRepository(mongoClient *client_mongo.HotelMongoClient, config *config.Configuration) interfaces.ICompanyRepository {
	return &CompanyRepository{
		MongoClient: mongoClient,
		Config:      config,
	}
}
