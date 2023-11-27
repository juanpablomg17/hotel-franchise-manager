package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/flexuxs/clubHubApp/internal/app"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.Initialize(ctx, r)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cancel()
		// Add any additional cleanup/shutdown logic here
		os.Exit(1)
	}()

	// Run the application
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
