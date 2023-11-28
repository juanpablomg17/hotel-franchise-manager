package model

type CompanyAggregate struct {
	Id      string  `bson:"id" json:"id"`
	Company Company `bson:"company" json:"company"`
}
