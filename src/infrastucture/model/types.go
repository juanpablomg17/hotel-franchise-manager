package infra_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Contact struct {
	Email      string             `json:"email" bson:"email"`
	Phone      string             `json:"phone" bson:"phone"`
	LocationId primitive.ObjectID `json:"locationId" bson:"locationId"`
}

type Owner struct {
	FirstName string  `json:"first_name" bson:"first_name"`
	LastName  string  `json:"last_name" bson:"last_name"`
	Contact   Contact `json:"contact" bson:"contact"`
}

type Information struct {
	Name       string             `json:"name" bson:"name"`
	TaxNumber  string             `json:"tax_number" bson:"tax_number"`
	LocationId primitive.ObjectID `json:"locationId" bson:"locationId"`
}

type Franchise struct {
	Name       string             `json:"name" bson:"name"`
	URL        string             `json:"url" bson:"url"`
	LocationId primitive.ObjectID `json:"locationId" bson:"locationId"`
}

type CompanyModel struct {
	Id          string      `json:"id" bson:"id"`
	Owner       Owner       `json:"owner" bson:"owner"`
	Information Information `json:"informacion" bson:"informacion"`
	Franchises  []Franchise `json:"franchises" bson:"franchises"`
}
