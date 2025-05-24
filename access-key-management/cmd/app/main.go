package main

import (
	"access-key-management/cmd/routes"
	"access-key-management/internal/database"
	"access-key-management/internal/stream"
	"access-key-management/pkg/middlewares"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func initApp() error {
	return errors.Join(stream.GetStreamer().Ping(), database.GetDb().Ping(), middlewares.InitLogger())
}

func main() {
	err := initApp()
	if err != nil {
		panic(err)
	}
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
