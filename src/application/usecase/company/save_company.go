package company

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/flexuxs/clubHubApp/src/domain/company/model"
	franchiseTypes "github.com/flexuxs/clubHubApp/src/domain/franchise/model"
	providers "github.com/flexuxs/clubHubApp/src/domain/providers/network"
	"github.com/flexuxs/clubHubApp/src/domain/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cu *CompanyUseCases) SaveCompany(ctx context.Context, company *model.CompanyAggregate) *utils.GenericCommandResponse {

	cu.ManageNetworkInfo(company)

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

func (cu *CompanyUseCases) ManageNetworkInfo(company *model.CompanyAggregate) {
	var wg sync.WaitGroup

	for _, franchise := range company.Company.Franchises {
		wg.Add(1)
		go func(franchise franchiseTypes.Franchise) {
			defer wg.Done()
			sslInfo, err := cu.siteFetcherService.GetSSLInfo(franchise.URL)
			if err != nil {
				fmt.Printf("Error fetching SSL info for %s: %v\n", franchise.URL, err)
				return
			}

			prinftInfo := utils.PrintStructInfoWithoutZeroValues(sslInfo)
			fmt.Printf("SSL info for %s: %v\n", franchise.URL, prinftInfo)
			fmt.Println("-----------------------------------")
			fmt.Printf("Start time: %v\n", sslInfo.StartTime)
			fmt.Printf("Test time: %v\n", sslInfo.TestTime)
			fmt.Printf("Status: %v\n", sslInfo.Status)
			fmt.Printf("Protocol: %v\n", sslInfo.Protocol)
			cu.showDomainInfo(sslInfo.Endpoints)
		}(franchise)
	}

	wg.Wait()
}

func (cu *CompanyUseCases) showDomainInfo(servers []providers.Endpoint) {
	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go func(server providers.Endpoint) {
			defer wg.Done()
			domainInfo, err := cu.siteFetcherService.GetWhoisInfo(server.IPAddress)
			if err != nil {
				fmt.Printf("Error fetching domain info for %s: %v\n", server.IPAddress, err)
				return
			}
			fmt.Printf("Domain info for %s: %v\n", server.IPAddress, domainInfo)
		}(server)
	}

	wg.Wait()
}
