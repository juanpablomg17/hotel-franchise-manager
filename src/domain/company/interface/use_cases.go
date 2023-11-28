package interfaces

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
)

type ICompanyUseCase interface {
	SaveCompany(ctx context.Context, company *model.CompanyAggregate) *utils.GenericCommandResponse
	UpdateCompany(ctx context.Context, company *model.CompanyAggregate) error
	GetCompany(ctx context.Context, filterRequest *model.Filterrequest) ([]*model.Company, error)
}
