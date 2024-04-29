package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	db "github.com/diegom0ta/go-mongodb/internal/database"
	"github.com/diegom0ta/go-mongodb/internal/http/server"
)

var ctx context.Context

func main() {
	db.Connect(ctx)

	go func() {
		server.Run(3333)
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	log.Println("Graceful shutdown started...")

	server.Shutdown()

	db.Disconnect(ctx)

	log.Println("Server is down")
}
