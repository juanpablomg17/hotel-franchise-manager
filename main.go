package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/flexuxs/clubHubApp/docs"

	"github.com/flexuxs/clubHubApp/src/api/fxmodule"
)

// @title ClubHub API
// @version	1.0
// @description Hoteld app for club management

// @host 	localhost:8888
// @BasePath /api
func main() {
	app := fxmodule.NewApp()

	go func() {
		if err := app.Start(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err)
	}

	app.Run()
}
