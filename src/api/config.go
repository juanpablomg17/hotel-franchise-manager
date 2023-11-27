package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/flexuxs/clubHubApp/src/api/routes"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	return gin.Default()
}

func RegisterRoutes(engine *gin.Engine, userHandler *routes.UserService) {
	router := engine.Group("/api/v1")
	router.GET("/user", userHandler.GetUser)
}

func StartServer(lifecycle fx.Lifecycle, engine *gin.Engine) {
	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			return server.Shutdown(ctx)
		},
	})
}