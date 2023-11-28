package interfaces

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
)

type ICompanyRepository interface {
	Save(ctx context.Context, company *model.CompanyAggregate) error
	Update(ctx context.Context, company *model.CompanyAggregate) error
}

type ICompanyFinder interface {
	GetCompanyByFilter(ctx context.Context, filterRequest *model.Filterrequest) ([]infra_model.CompanyModel, error)
}
