package infra_providers

import (
	"encoding/json"
	"net/http"
	"time"

	providers "github.com/flexuxs/clubHubApp/src/domain/providers/network"
	"github.com/gocolly/colly"
	"github.com/likexian/whois"
)

// GetImageUrl() (string, error)
// IsWebsiteAvailable() (bool, error)
// GetCommunicationProtocol() (string, error)
// GetHopCountAndServers() ([]string, error)
// GetDomainInfo() (DomainInfo, error)

const maxRetries = 3
const delay = 2 * time.Second

func (si *SiteInfoFetcherService) GetImageUrl(protocol, host string) (string, error) {
	var imageURL string

	// Create a new collector
	c := colly.NewCollector()

	// Define behavior when an image is found
	c.OnHTML("img", func(e *colly.HTMLElement) {
		imageURL = e.Attr("src")
	})

	// Visit the provided URL
	url := protocol + "://" + host
	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}
func (si *SiteInfoFetcherService) IsWebsiteAvailable() (bool, error) {
	return true, nil
}

func (si *SiteInfoFetcherService) GetCommunicationProtocol() (string, error) {
	return "", nil
}

func (si *SiteInfoFetcherService) GetHopCountAndServers() ([]string, error) {
	return []string{}, nil
}

func (si *SiteInfoFetcherService) GetDomainInfo() (providers.DomainInfo, error) {
	return providers.DomainInfo{}, nil
}

func (si *SiteInfoFetcherService) GetSSLInfo(url string) (*providers.SSLLabResponse, error) {
	apiURL := "https://api.ssllabs.com/api/v3/analyze?host=" + url

	var resp *http.Response
	var err error

	for i := 0; i < maxRetries; i++ {
		resp, err = http.Get(apiURL)
		if err == nil {
			break
		}
		time.Sleep(delay)
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	sslInfo := &providers.SSLLabResponse{}
	err = json.NewDecoder(resp.Body).Decode(&sslInfo)
	if err != nil {
		return nil, err
	}
	return sslInfo, nil
}

func (si *SiteInfoFetcherService) GetWhoisInfo(url string, server ...string) (string, error) {
	result, err := whois.Whois(url, server...)
	if err != nil {
		return "", err
	}
	return result, nil
}
