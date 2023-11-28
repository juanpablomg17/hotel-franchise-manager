package providers

// type SiteInfoFetcher interface {
// 	ScrapeImageURL(url string) (string, error)
// 	IsWebsiteAvailable(url string) bool
// 	GetWhoisInfo(url string, server ...string) (string, error)
// 	GetSSLInfo(url string) (*SSLLabResponse, error)
// }

type ISiteInfoFetcher interface {
	GetImageUrl(protocol, host string) (string, error)
	IsWebsiteAvailable() (bool, error)
	GetCommunicationProtocol() (string, error)
	GetHopCountAndServers() ([]string, error)
	GetDomainInfo() (DomainInfo, error)
	GetSSLInfo(url string) (*SSLLabResponse, error)
	GetWhoisInfo(url string, server ...string) (string, error)
}
