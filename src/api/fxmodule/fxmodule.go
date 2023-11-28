package fxmodule

import (
	"github.com/flexuxs/clubHubApp/src/api"
	"github.com/flexuxs/clubHubApp/src/api/config"
	controller "github.com/flexuxs/clubHubApp/src/api/controller/company"
	client_mongo "github.com/flexuxs/clubHubApp/src/api/fxmodule/client"
	"github.com/flexuxs/clubHubApp/src/api/routes"
	companyApplication "github.com/flexuxs/clubHubApp/src/application/usecase/company"
	"github.com/flexuxs/clubHubApp/src/infrastucture/finder"
	infra_providers "github.com/flexuxs/clubHubApp/src/infrastucture/providers/network"
	"github.com/flexuxs/clubHubApp/src/infrastucture/repository"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		// config
		fx.Provide(config.ProvideConfiguration),

		// infrastructure
		fx.Provide(client_mongo.NewHotelMongoClient),
		fx.Provide(repository.NewCompanyRepository),
		fx.Provide(finder.NewCompanyFinder),
		fx.Provide(infra_providers.NewSiteInfoFetcher),

		// application
		fx.Provide(companyApplication.NewCompanyUseCases),

		// controller
		fx.Provide(controller.NewController),

		// api
		fx.Provide(api.NewGinEngine),
		fx.Provide(routes.NewCompanyRoutes),

		// register routes
		fx.Invoke(api.RegisterRoutes),

		// start server
		fx.Invoke(api.StartServer),
	)
}
