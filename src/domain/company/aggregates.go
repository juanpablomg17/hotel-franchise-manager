package company

import "github.com/flexuxs/clubHubApp/src/domain/company/model"

type CompanyAggregate struct {
	Id      string        `bson:"_id" json:"id"`
	Company model.Company `bson:"company" json:"company"`
}
