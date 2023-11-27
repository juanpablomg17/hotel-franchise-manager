package model

type CompanyAggregate struct {
	Id      string  `bson:"_id" json:"id"`
	Company Company `bson:"company" json:"company"`
}
