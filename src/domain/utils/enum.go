package utils

import (
	infra_model "github.com/flexuxs/clubHubApp/src/infrastucture/model"
)

type GenericCommandResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type GetCompaniesResponse struct {
	StatusCode int                        `json:"status_code"`
	Message    string                     `json:"message"`
	Companies  []infra_model.CompanyModel `json:"companies"`
}
