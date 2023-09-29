package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
	DeletedAt string `json:"-"`
}

type Products []*Product

func AllProducts() Products {
	return Products(products)
}

func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func FromJSON(r io.Reader) (*Product, error) {
	product := &Product{ CreatedAt: time.Now().UTC().String(), UpdatedAt: time.Now().UTC().String()}
	err := json.NewDecoder(r).Decode(product)
	fmt.Printf("**json: %v\n", product)
	return product, err
}

func AddProduct(p *Product) {
	products = append(products, p)
}

var products = []*Product{
	&Product{
		ID: 12345,
		Name: "Coffee",
		Description: "A warm bean stained water drink",
		Price: 2.30,
		SKU: "abc123",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Product{
		ID: 12346,
		Name: "Esspresso",
		Description: "Like coffee, but smaller and more caffeine",
		Price: 1.99,
		SKU: "xyz123",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}