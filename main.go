package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/flexuxs/clubHubApp/src/api/fxmodule"
)

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
