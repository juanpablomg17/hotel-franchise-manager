package infra_providers

import (
	"github.com/flexuxs/clubHubApp/src/api/config"
	providers "github.com/flexuxs/clubHubApp/src/domain/providers/network"
)

type SiteInfoFetcherService struct {
	config *config.Configuration
}

func NewSiteInfoFetcher(config *config.Configuration) providers.ISiteInfoFetcher {
	return &SiteInfoFetcherService{
		config: config,
	}
}
