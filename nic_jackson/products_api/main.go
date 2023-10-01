package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/richard12511/product_api/handlers"
)

func main(){
	logger := log.New(os.Stdout, "products_api", log.LstdFlags)
	ph := handlers.NewProductsHandler(logger)

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	router.Handle("/products", ph)
	router.Handle("/products/{key}", ph)

	server := &http.Server{
		Addr: ":9090",
		Handler: router,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	sigChan := make(chan os.Signal, 2)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, syscall.SIGTERM)

	go func() { server.ListenAndServe() }()

	<- sigChan

	logger.Println("Graceful shutdown starting now")
	ctx, _ := context.WithTimeout(context.Background(), 90 * time.Second)
	server.Shutdown(ctx)
}