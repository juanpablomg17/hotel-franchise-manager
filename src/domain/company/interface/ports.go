package interfaces

import (
	"context"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
)

type ICompanyRepository interface {
	Save(ctx context.Context, company *model.CompanyAggregate) error
	Update(ctx context.Context, company *model.CompanyAggregate) error
}

type ICompanyFinder interface {
	Get(ctx context.Context, filterRequest *model.Filterrequest) ([]model.Company, error)
}
