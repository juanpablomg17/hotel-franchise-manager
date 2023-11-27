package interfaces

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
)

type ICompanyUseCase interface {
	SaveCompany(ctx context.Context, company *model.CompanyAggregate) error
	UpdateCompany(ctx context.Context, company *model.CompanyAggregate) error
	GetCompany(ctx context.Context, filterRequest *model.Filterrequest) ([]*model.Company, error)
}
