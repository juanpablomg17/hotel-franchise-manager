package repository

import (
	"context"
	"errors"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	locationTypes "github.com/flexuxs/clubHubApp/src/domain/location/model"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
	infra_services "github.com/flexuxs/clubHubApp/src/infrastucture/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *CompanyRepository) Save(ctx context.Context, company *model.CompanyAggregate) error {

	companyCollection, err := c.MongoClient.Client.Database(c.Config.Mongo.Database).Collection(c.Config.Mongo.CompanyCollection).Clone()

	companyModelToSave := &infra_model.CompanyModel{}

	companyModelToSave = infra_services.FromDomainCompanyToDatabaseModel(company)

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

	franchiseLocationIDs := make([]primitive.ObjectID, len(company.Company.Franchises))
	for i, franchise := range company.Company.Franchises {
		franchiseLocationID, err := c.saveLocation(ctx, &franchise.Location)
		if err != nil {
			return err
		}
		franchiseLocationIDs[i] = franchiseLocationID
	}

	companyModelToSave.Owner.Contact.LocationId = ownerLocationID
	companyModelToSave.Information.LocationId = infoLocationID

	for i := range companyModelToSave.Franchises {
		companyModelToSave.Franchises[i].LocationId = franchiseLocationIDs[i]
	}

	_, err = companyCollection.InsertOne(ctx, companyModelToSave)

	return err
}

func (c *CompanyRepository) saveLocation(ctx context.Context, location *locationTypes.Location) (primitive.ObjectID, error) {
	collection := c.MongoClient.Client.Database(c.Config.Mongo.Database).Collection(c.Config.Mongo.LocationCollection)

	result, err := collection.InsertOne(ctx, location)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to convert InsertedID to primitive.ObjectID")
	}

	return insertedID, nil
}

func (c *CompanyRepository) Update(ctx context.Context, company *infra_model.CompanyModel) error {
	companyCollection := c.MongoClient.Client.Database(c.Config.Mongo.Database).Collection(c.Config.Mongo.CompanyCollection)

	// Check if company exists by querying for the company ID
	filter := bson.M{"_id": company.Id}
	foundCompany := &infra_model.CompanyModel{}
	err := companyCollection.FindOne(ctx, filter).Decode(foundCompany)
	if err != nil {
		return err
	}

	// Update the company
	update := bson.M{
		"$set": bson.M{
			"franchises":  company.Franchises,
			"owner":       company.Owner,
			"informacion": company.Information,
		},
	}
	_, err = companyCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
