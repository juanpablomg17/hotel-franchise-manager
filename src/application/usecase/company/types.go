package company

import (
	interfaces "github.com/flexuxs/clubHubApp/src/domain/company/interface"
	"github.com/sirupsen/logrus"
)

type CompanyUseCases struct {
	Logger            *logrus.Logger
	CompanyFinder     interfaces.ICompanyFinder
	CompanyRepository interfaces.ICompanyRepository
}

func NewCompanyUseCases(logger *logrus.Logger, companyFinder interfaces.ICompanyFinder, companyRepository interfaces.ICompanyRepository) interfaces.ICompanyUseCase {
	return &CompanyUseCases{
		Logger:            logger,
		CompanyFinder:     companyFinder,
		CompanyRepository: companyRepository,
	}
}
