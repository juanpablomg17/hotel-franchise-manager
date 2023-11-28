package company

import (
	"context"
	"fmt"
	"net/http"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
)

func (cu *CompanyUseCases) GetCompany(ctx context.Context, filterRequest *model.Filterrequest) *utils.GetCompaniesResponse {

	response := &utils.GetCompaniesResponse{}

	companies, err := cu.CompanyFinder.GetCompanyByFilter(ctx, filterRequest)

	if err != nil {
		fmt.Println("holasss", err)
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Error getting companies"
		return response
	}

	response.StatusCode = http.StatusOK
	response.Message = "Companies retrieved successfully"
	response.Companies = companies

	return response
}
