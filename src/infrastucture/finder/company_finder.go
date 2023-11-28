package finder

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (cf *CompanyFinder) GetCompanyByFilter(ctx context.Context, filterRequest *model.Filterrequest) ([]infra_model.CompanyModel, error) {
	collection, err := cf.MongoClient.Client.Database(cf.Config.Mongo.Database).Collection(cf.Config.Mongo.CompanyCollection).Clone()
	if err != nil {
		return nil, err
	}

	filter := bson.M{}

	if filterRequest.CompanyId != "" {
		filter["id"] = filterRequest.CompanyId
	}

	if filterRequest.FranchiseName != "" {
		filter["franchises.name"] = filterRequest.FranchiseName
	}

	if filterRequest.CompanyName != "" {
		filter["informacion.name"] = filterRequest.CompanyName
	}

	if filterRequest.OwnerFirstName != "" {
		filter["$or"] = []bson.M{
			{"owner.first_name": filterRequest.OwnerFirstName},
		}
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var companies []infra_model.CompanyModel
	for cursor.Next(context.Background()) {
		var company infra_model.CompanyModel
		if err := cursor.Decode(&company); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	return companies, nil
}
