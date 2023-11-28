package company

import (
	interfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
)

type CompanyUseCases struct {
	CompanyFinder     interfaces.ICompanyFinder
	CompanyRepository interfaces.ICompanyRepository
}

func NewCompanyUseCases(companyFinder interfaces.ICompanyFinder, companyRepository interfaces.ICompanyRepository) interfaces.ICompanyUseCase {
	return &CompanyUseCases{
		CompanyFinder:     companyFinder,
		CompanyRepository: companyRepository,
	}
}
