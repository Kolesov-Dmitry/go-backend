package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"file-srv/internal/srv"
)

// go build -o .\bin\ .\cmd\file-srv\
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	fileServer := srv.NewFileServer()

	go fileServer.Start()
	log.Println("File server is started")

	// Handle ctrl+c
	<-ctx.Done()
	cancel()

	// Shutdown server
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	fileServer.Shutdown(ctx)
	cancel()
}
