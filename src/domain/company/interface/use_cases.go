package interfaces

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
)

type ICompanyUseCase interface {
	SaveCompany(ctx context.Context, company *model.CompanyAggregate) *utils.GenericCommandResponse
	UpdateCompany(ctx context.Context, company *infra_model.CompanyModel) *utils.GenericCommandResponse
	GetCompany(ctx context.Context, filterRequest *model.Filterrequest) *utils.GetCompaniesResponse
}
