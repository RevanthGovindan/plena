package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web3-tokeninfo/cmd/routes"
	"web3-tokeninfo/internal/database"
	"web3-tokeninfo/internal/services"
	"web3-tokeninfo/internal/stream"
	"web3-tokeninfo/pkg/utils"

	"github.com/gorilla/mux"
)

func initApp() error {
	var err = errors.Join(stream.GetStreamer().Ping(), database.GetDb().Ping())
	if err == nil {
		go stream.GetStreamer().Subscribe(utils.SUBSCRIBE_TOPIC, func(msg string) {
			services.HandleEvents(msg)
		})
	}
	return err
}

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	routes := routes.GetRoutes(router)
	server := &http.Server{
		Addr:         ":8082",
		Handler:      routes,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// Run server in a goroutine so we can shut it down gracefully
	go func() {
		log.Println("Server started on :8082")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	<-stop
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited gracefully")
}
