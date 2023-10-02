package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/richard12511/product_api/data"
)

type Products struct {
	logger *log.Logger
}

func NewProductsHandler(logger *log.Logger) *Products {
	return &Products{logger}
}

func (p *Products) GetProducts(rw http.ResponseWriter, req *http.Request){
	products := data.AllProducts()
	err := products.ToJSON(rw)

	if err != nil {
		p.logger.Fatalf("Error retrieving products: %s", err.Error())
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
}

func (p *Products) PostProduct(rw http.ResponseWriter, req *http.Request){
	product, err := data.FromJSON(req.Body)

	if err != nil {
		p.logger.Fatalf("Error creating product: %s", err.Error())
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}
	data.AddProduct(product)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	key, err := strconv.Atoi(vars["key"])
	if err != nil {
		http.Error(rw, "key is not integer", http.StatusBadRequest)
	}

	product, err := data.FromJSON(req.Body)
	if err != nil {
		http.Error(rw, "Bad Request", http.StatusBadRequest)
	}

	data.UpdateProduct(key, product)	
}

//curl localhost:9090/products -d '{ "id":239, "name":"Tea", "description": "Nice warm tea", "price":1.20, "sku":"abba"}'
//curl localhost:9090 | jq
//curl localhost:9090 -d { "id":6, "name":"Tea", "description": "Nice warm tea", "sku":"abba"}'
//curl localhost:9090 -d '{ "id": 1, "name": "Tea", "description": "a nice cup of tea" }' | jq
//curl localhost:9090/products -d 
//curl localhost:9090/products/1 -XPUT -d '{"name": "tea", "description": "a nice cup of tea, yea buddy!"}' | jq
