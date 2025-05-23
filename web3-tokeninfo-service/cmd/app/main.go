package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web3-tokeninfo/cmd/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes := routes.GetRoutes(router)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	// Run server in a goroutine so we can shut it down gracefully
	go func() {
		log.Println("Server started on :8080")
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
