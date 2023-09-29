package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/richard12511/product_api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProductsHandler(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	products := data.AllProducts()
	err := json.NewEncoder(rw).Encode(products)

	if err != nil {
		p.logger.Fatalf("Error retrieving products: %s", err.Error())
	}
}
