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

func main() {
	ctx := context.Background()

	db.Connect(ctx)

	go func() {
		server.Run(3333)
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	log.Println("Graceful shutdown started...")

	db.Disconnect(ctx)

	server.Shutdown()

	log.Println("Server is down")
}
