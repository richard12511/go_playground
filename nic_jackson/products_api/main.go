package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/richard12511/product_api/handlers"
)

func main(){
	logger := log.New(os.Stdout, "products_api", log.LstdFlags)
	ph := handlers.NewProductsHandler(logger)

	mux := http.NewServeMux()
	mux.Handle("/products", ph)

	server := &http.Server{
		Addr: ":9090",
		Handler: mux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	server.ListenAndServe()
}