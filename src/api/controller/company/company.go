package controller

import (
	"context"

	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/flexuxs/clubHubApp/src/api/mapper"
	companyInterfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
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

// @Summary Save a new company
// @Description SaveCompany saves a new company using the provided DTO.
// @Accept json
// @Produce json
// @Param company body dto.CompanyDTO true "Company data"
// @Success 200 {object} utils.GenericCommandResponse
// @Router /company [post]
func (c *Controller) SaveCompany(company *dto.CompanyDTO) *utils.GenericCommandResponse {

	companyAggregate := mapper.FromCompanyDtoToCompanyAggregate(company)

	response := c.CompanyUseCases.SaveCompany(context.Background(), &companyAggregate)

	return response
}

// @Summary Get company by filter
// @Description GetCompanyByfilter retrieves companies based on the provided filter.
// @Accept json
// @Produce json
// @Param companyFilter query dto.CompanyFilterDTO true "Filter criteria"
// @Success 200 {object} utils.GetCompaniesResponse
// @Failure 500 {object} error
// @Router /company [get]
func (c *Controller) GetCompanyByfilter(companyFilter *dto.CompanyFilterDTO) (*utils.GetCompaniesResponse, error) {

	getCompantRquest := &model.Filterrequest{}
	err := utils.AdapStructs(getCompantRquest, companyFilter)

	if err != nil {
		return nil, err
	}

	response := c.CompanyUseCases.GetCompany(context.Background(), getCompantRquest)

	return response, nil
}

// @Summary Update company
// @Description UpdateCompany updates an existing company using the provided model.
// @Accept json
// @Produce json
// @Param id path string true "Company ID"
// @Param company body infra_model.CompanyModel true "Updated company data"
// @Success 200 {object} utils.GenericCommandResponse
// @Router /company/{id} [put]
func (c *Controller) UpdateCompany(companyDto *infra_model.CompanyModel) *utils.GenericCommandResponse {

	response := c.CompanyUseCases.UpdateCompany(context.Background(), companyDto)

	return response
}
