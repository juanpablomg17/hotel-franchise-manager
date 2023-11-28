package company

import (
	"context"
	"net/http"

	company_service "github.com/flexuxs/clubHubApp/src/domain/company/services"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
)

func (cu *CompanyUseCases) UpdateCompany(ctx context.Context, company *infra_model.CompanyModel) *utils.GenericCommandResponse {

	response := &utils.GenericCommandResponse{}

	foundCompany, err := cu.CompanyFinder.GetCompanyById(ctx, company.Id)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Error getting companies"
	}

	if foundCompany == nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Company not found"
	}

	hasChanges, err := company_service.HasChanges(foundCompany, company)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Error getting companies"
	}

	if !hasChanges {
		response.StatusCode = http.StatusBadRequest
		response.Message = "No changes detected"
		return nil
	}

	err = cu.CompanyRepository.Update(ctx, company)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Error updating company"
		return response
	}

	response.StatusCode = http.StatusOK
	response.Message = "Company updated successfully"

	return response
}
