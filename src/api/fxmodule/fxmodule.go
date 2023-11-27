package fxmodule

import (
	"github.com/flexuxs/clubHubApp/src/api"
	"github.com/flexuxs/clubHubApp/src/api/routes"
	infrastructure "github.com/flexuxs/clubHubApp/src/infrastucture"
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(api.NewGinEngine),
		fx.Provide(routes.NewUserService),
		fx.Provide(infrastructure.NewUserRepository),
		fx.Invoke(api.RegisterRoutes),
		fx.Invoke(api.StartServer),
	)
}
