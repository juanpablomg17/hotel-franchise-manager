package repository

import "github.com/flexuxs/clubHubApp/src/domain/company/model"

func FromDomainCompanyToDatabaseModel(company *model.CompanyAggregate) *CompanyModel {
	companyModel := &CompanyModel{}

	// Map Owner
	companyModel.Company.Owner.FirstName = company.Company.Owner.FirstName
	companyModel.Company.Owner.LastName = company.Company.Owner.LastName
	companyModel.Company.Owner.Contact.Email = company.Company.Owner.Contact.Email
	companyModel.Company.Owner.Contact.Phone = company.Company.Owner.Contact.Phone
	companyModel.Company.Owner.Contact.LocationId = company.Company.Owner.Contact.Location.Id

	// Map Information
	companyModel.Company.Information.Name = company.Company.Information.Name
	companyModel.Company.Information.TaxNumber = company.Company.Information.TaxNumber
	companyModel.Company.Information.LocationId = company.Company.Information.Location.Id

	// Map Franchises

	companyModel.Company.Franchises = make([]Franchise, len(company.Company.Franchises))
	for i, franchise := range company.Company.Franchises {
		companyModel.Company.Franchises[i].Name = franchise.Name
		companyModel.Company.Franchises[i].URL = franchise.URL
		companyModel.Company.Franchises[i].LocationId = franchise.Location.Id
	}

	return companyModel
}
