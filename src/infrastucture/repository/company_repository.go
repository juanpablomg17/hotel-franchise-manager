package repository

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	locationTypes "github.com/flexuxs/clubHubApp/src/domain/location/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *CompanyRepository) Save(ctx context.Context, company *model.CompanyAggregate) error {

	companyCollection, err := c.MongoClient.Client.Database(c.Config.Mongo.Database).Collection(c.Config.Mongo.CompanyCollection).Clone()

	companyModelToSave := &CompanyModel{}

	companyModelToSave = FromDomainCompanyToDatabaseModel(company)

	if err != nil {
		return err
	}

	ownerLocationID, err := c.saveLocation(ctx, &company.Company.Owner.Contact.Location)
	if err != nil {
		return err
	}

	infoLocationID, err := c.saveLocation(ctx, &company.Company.Information.Location)
	if err != nil {
		return err
	}

	franchiseLocationIDs := make([]string, len(company.Company.Franchises))
	for i, franchise := range company.Company.Franchises {
		franchiseLocationID, err := c.saveLocation(ctx, &franchise.Location)
		if err != nil {
			return err
		}
		franchiseLocationIDs[i] = franchiseLocationID
	}

	companyModelToSave.Company.Owner.Contact.LocationId = ownerLocationID
	companyModelToSave.Company.Information.LocationId = infoLocationID

	for i, franchise := range companyModelToSave.Company.Franchises {
		franchise.LocationId = franchiseLocationIDs[i]
	}

	_, err = companyCollection.InsertOne(ctx, company)

	return err
}

func (c *CompanyRepository) saveLocation(ctx context.Context, location *locationTypes.Location) (string, error) {
	collection, err := c.MongoClient.Client.Database(c.Config.Mongo.Database).Collection(c.Config.Mongo.LocationCollection).Clone()

	if err != nil {
		return "", err
	}

	result, err := collection.InsertOne(ctx, location)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (c *CompanyRepository) Update(ctx context.Context, company *model.CompanyAggregate) error {
	return nil
}
