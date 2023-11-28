package company

import (
	"context"
	"net/http"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cu *CompanyUseCases) SaveCompany(ctx context.Context, company *model.CompanyAggregate) *utils.GenericCommandResponse {
	response := &utils.GenericCommandResponse{}
	err := cu.CompanyRepository.Save(ctx, company)

	if mongo.IsDuplicateKeyError(err) {
		response.StatusCode = http.StatusConflict
		response.Message = "Company already exists"
		return response
	}

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Message = "Error saving company"
		return response
	}
	response.StatusCode = http.StatusCreated
	response.Message = "Company saved successfully"
	return response
}
