package company

import (
	interfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
	providers "github.com/flexuxs/clubHubApp/src/domain/providers/network"
)

type CompanyUseCases struct {
	CompanyFinder      interfaces.ICompanyFinder
	CompanyRepository  interfaces.ICompanyRepository
	siteFetcherService providers.ISiteInfoFetcher
}

func NewCompanyUseCases(companyFinder interfaces.ICompanyFinder, companyRepository interfaces.ICompanyRepository, siteFetcherService providers.ISiteInfoFetcher) interfaces.ICompanyUseCase {
	return &CompanyUseCases{
		CompanyFinder:      companyFinder,
		CompanyRepository:  companyRepository,
		siteFetcherService: siteFetcherService,
	}
}
