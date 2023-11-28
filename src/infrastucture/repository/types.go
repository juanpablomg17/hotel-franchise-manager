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

type Contact struct {
	Email      string `json:"email" bson:"email"`
	Phone      string `json:"phone" bson:"phone"`
	LocationId string `json:"locationId" bson:"locationId"`
}

type Owner struct {
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Contact   Contact `json:"contact" bson:"contact"`
}

type Information struct {
	Name       string `json:"name" bson:"name"`
	TaxNumber  string `json:"tax_number" bson:"tax_number"`
	LocationId string `json:"locationId" bson:"locationId"`
}

type Franchise struct {
	Name       string `json:"name" bson:"name"`
	URL        string `json:"url" bson:"url"`
	LocationId string `json:"locationId" bson:"locationId"`
}

type Company struct {
	Owner       Owner       `json:"owner" bson:"owner"`
	Information Information `json:"informacion" bson:"informacion"`
	Franchises  []Franchise `json:"franchises" bson:"franchises"`
}

type CompanyModel struct {
	Company Company `json:"company" bson:"company"`
}
