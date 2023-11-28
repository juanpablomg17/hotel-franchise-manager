package controller

import (
	"context"

	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/flexuxs/clubHubApp/src/api/mapper"
	companyInterfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
)

type Controller struct {
	CompanyUseCases companyInterfaces.ICompanyUseCase
}

// NewController initializes a new Controller struct.
func NewController(companyUseCases companyInterfaces.ICompanyUseCase) *Controller {
	return &Controller{
		CompanyUseCases: companyUseCases,
	}
}

func (c *Controller) SaveCompany(company *dto.CompanyDTO) *utils.GenericCommandResponse {

	companyAggregate := mapper.FromCompanyDtoToCompanyAggregate(company)

	response := c.CompanyUseCases.SaveCompany(context.Background(), &companyAggregate)

	return response
}

func (c *Controller) GetCompanyByfilter(companyFilter *dto.CompanyFilterDTO) (*utils.GetCompaniesResponse, error) {

	getCompantRquest := &model.Filterrequest{}
	err := utils.AdapStructs(getCompantRquest, companyFilter)

	if err != nil {
		return nil, err
	}

	response := c.CompanyUseCases.GetCompany(context.Background(), getCompantRquest)

	return response, nil
}
