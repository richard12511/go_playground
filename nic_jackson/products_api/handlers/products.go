package handlers

import (
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
	if req.Method == http.MethodGet {
		p.getProducts(rw)
	}

	if req.Method == http.MethodPost {
		p.postProduct(rw, req)
	}
}

func (p *Products) getProducts(rw http.ResponseWriter){
	products := data.AllProducts()
	err := products.ToJSON(rw)

	if err != nil {
		p.logger.Fatalf("Error retrieving products: %s", err.Error())
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
}

func (p *Products) postProduct(rw http.ResponseWriter, req *http.Request){
	product, err := data.FromJSON(req.Body)

	if err != nil {
		p.logger.Fatalf("Error creating product: %s", err.Error())
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

//curl localhost:9090/products -d '{ "id":239, "name":"Tea", "description": "Nice warm tea", "price":1.20, "sku":"abba"}'
//curl localhost:9090 | jq
//curl localhost:9090 -d { "id":6, "name":"Tea", "description": "Nice warm tea", "sku":"abba"}'
//curl localhost:9090 -d '{ "id": 1, "name": "Tea", "description": "a nice cup of tea" }' | jq
//curl localhost:9090/products -d 
