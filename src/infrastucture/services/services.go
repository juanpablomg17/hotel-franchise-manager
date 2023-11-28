package infra_services

import (
	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	franchise_model "github.com/flexuxs/clubHubApp/src/domain/franchise/model"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
)

func FromDomainCompanyToDatabaseModel(company *model.CompanyAggregate) *infra_model.CompanyModel {
	companyModel := &infra_model.CompanyModel{}

	// Map Owner
	companyModel.Id = company.Id
	companyModel.Owner.FirstName = company.Company.Owner.FirstName
	companyModel.Owner.LastName = company.Company.Owner.LastName
	companyModel.Owner.Contact.Email = company.Company.Owner.Contact.Email
	companyModel.Owner.Contact.Phone = company.Company.Owner.Contact.Phone

	// Map Information
	companyModel.Information.Name = company.Company.Information.Name
	companyModel.Information.TaxNumber = company.Company.Information.TaxNumber

	// Map Franchises

	companyModel.Franchises = make([]infra_model.Franchise, len(company.Company.Franchises))
	for i, franchise := range company.Company.Franchises {
		companyModel.Franchises[i].Name = franchise.Name
		companyModel.Franchises[i].URL = franchise.URL
	}

	return companyModel
}

func FromDatabaseCompanyToDomainModelArray(companies []infra_model.CompanyModel) []model.Company {
	companyAggregates := make([]model.Company, len(companies))

	for i, company := range companies {
		companyDomain := &model.Company{}

		// Map Owner
		companyDomain.Id = company.Id
		companyDomain.Owner.FirstName = company.Owner.FirstName
		companyDomain.Owner.LastName = company.Owner.LastName
		companyDomain.Owner.Contact.Email = company.Owner.Contact.Email
		companyDomain.Owner.Contact.Phone = company.Owner.Contact.Phone

		// Map Information
		companyDomain.Information.Name = company.Information.Name
		companyDomain.Information.TaxNumber = company.Information.TaxNumber

		// Map Franchises
		companyDomain.Franchises = make([]franchise_model.Franchise, len(company.Franchises))
		for j, franchise := range company.Franchises {
			companyDomain.Franchises[j].Name = franchise.Name
			companyDomain.Franchises[j].URL = franchise.URL
		}

		companyAggregates[i] = *companyDomain
	}

	return companyAggregates
}
