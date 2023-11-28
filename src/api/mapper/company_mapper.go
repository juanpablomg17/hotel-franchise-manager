package mapper

import (
	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	franchiseTypes "github.com/flexuxs/clubHubApp/src/domain/franchise/model"
	locantionTypes "github.com/flexuxs/clubHubApp/src/domain/location/model"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
	"github.com/google/uuid"
)

func FromCompanyDtoToCompanyAggregate(companyDto *dto.CompanyDTO) model.CompanyAggregate {
	company := model.Company{
		Owner: model.Owner{
			FirstName: companyDto.Company.Owner.FirstName,
			LastName:  companyDto.Company.Owner.LastName,
			Contact: model.ContactComany{
				Email: companyDto.Company.Owner.Contact.Email,
				Phone: companyDto.Company.Owner.Contact.Phone,
				Location: locantionTypes.Location{
					City:    companyDto.Company.Owner.Contact.Location.City,
					Country: companyDto.Company.Owner.Contact.Location.Country,
					Address: companyDto.Company.Owner.Contact.Location.Address,
					ZipCode: companyDto.Company.Owner.Contact.Location.ZipCode,
				},
			},
		},
		Information: model.Information{
			Name:      companyDto.Company.Information.Name,
			TaxNumber: companyDto.Company.Information.TaxNumber,
			Location: locantionTypes.Location{
				City:    companyDto.Company.Information.Location.City,
				Country: companyDto.Company.Information.Location.Country,
				Address: companyDto.Company.Information.Location.Address,
				ZipCode: companyDto.Company.Information.Location.ZipCode,
			},
		},
		Franchises: make([]franchiseTypes.Franchise, 0, len(companyDto.Company.Franchises)),
	}

	for _, franchiseDto := range companyDto.Company.Franchises {
		franchise := franchiseTypes.Franchise{
			Name: franchiseDto.Name,
			URL:  franchiseDto.URL,
			Location: locantionTypes.Location{
				City:    franchiseDto.Location.City,
				Country: franchiseDto.Location.Country,
				Address: franchiseDto.Location.Address,
				ZipCode: franchiseDto.Location.ZipCode,
			},
		}
		company.Franchises = append(company.Franchises, franchise)
	}

	return model.CompanyAggregate{
		Id:      uuid.New().String(),
		Company: company,
	}
}

func FromDomainCompanyToDatabaseModelArray(companies []model.Company) []infra_model.CompanyModel {
	companyModels := []infra_model.CompanyModel{}

	for _, company := range companies {
		for _, franchise := range company.Franchises {
			companyModel := &infra_model.CompanyModel{}

			// Map Owner
			companyModel.Id = company.Id
			companyModel.Owner.FirstName = company.Owner.FirstName
			companyModel.Owner.LastName = company.Owner.LastName
			companyModel.Owner.Contact.Email = company.Owner.Contact.Email
			companyModel.Owner.Contact.Phone = company.Owner.Contact.Phone

			// Map Information
			companyModel.Information.Name = company.Information.Name
			companyModel.Information.TaxNumber = company.Information.TaxNumber

			// Map Franchises
			companyModel.Franchises = make([]infra_model.Franchise, len(company.Franchises))
			companyModel.Franchises[0].Name = franchise.Name
			companyModel.Franchises[0].URL = franchise.URL

			companyModels = append(companyModels, *companyModel)
		}
	}

	return companyModels
}
