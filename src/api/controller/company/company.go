package controller

import (
	"context"

	dto "github.com/flexuxs/clubHubApp/src/api/dto/company"
	"github.com/flexuxs/clubHubApp/src/api/mapper"
	companyInterfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	CompanyUseCases companyInterfaces.ICompanyUseCase
	Logger          *logrus.Logger
}

// NewController initializes a new Controller struct.
func NewController(companyUseCases companyInterfaces.ICompanyUseCase, logger *logrus.Logger) *Controller {
	return &Controller{
		CompanyUseCases: companyUseCases,
		Logger:          logger,
	}
}

func (c *Controller) SaveCompany(company *dto.CompanyDTO) *utils.GenericCommandResponse {

	companyAggregate := mapper.FromCompanyDtoToCompanyAggregate(company)

	err := c.CompanyUseCases.SaveCompany(context.Background(), &companyAggregate)
	if err != nil {
		return &utils.GenericCommandResponse{
			StatusCode: 500,
			Message:    "Error saving company",
		}
	}

	return nil
}
