package model

import (
	franchiseTypes "github.com/flexuxs/clubHubApp/src/domain/franchise/model"
	locationTypes "github.com/flexuxs/clubHubApp/src/domain/location/model"
)

type Owner struct {
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Contact   ContactComany `json:"contact"`
}

type ContactComany struct {
	Email    string                 `json:"email"`
	Phone    string                 `json:"phone"`
	Location locationTypes.Location `json:"location"`
}

type Information struct {
	Name      string                 `json:"name"`
	TaxNumber string                 `json:"tax_number"`
	Location  locationTypes.Location `json:"location"`
}

type Company struct {
	Id          string                     `json:"id"`
	Owner       Owner                      `json:"owner"`
	Information Information                `json:"informacion"`
	Franchises  []franchiseTypes.Franchise `json:"franchises"`
}
