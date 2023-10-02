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

func AddProduct(p *Product) {
	p.ID = getNextID()
	products = append(products, p)
}

func UpdateProduct(id int, p *Product) error {
	_, i, err := FindProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	products[i] = p
	return nil
}

func FindProduct(id int) (*Product, int, error) {
	for i, p := range products {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, 0, ErrorProdNotFound
}

func (p *Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(p)
}

func FromJSON(r io.Reader) (*Product, error) {
	product := &Product{ CreatedAt: time.Now().UTC().String(), UpdatedAt: time.Now().UTC().String()}
	err := json.NewDecoder(r).Decode(product)
	return product, err
}

func getNextID() int {
	lp := products[len(products) - 1]
	return lp.ID + 1
}

var ErrorProdNotFound = fmt.Errorf("Product not found")

var products = []*Product{
	&Product{
		ID: 1,
		Name: "Coffee",
		Description: "A warm bean stained water drink",
		Price: 2.30,
		SKU: "abc123",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	&Product{
		ID: 2,
		Name: "Esspresso",
		Description: "Like coffee, but smaller and more caffeine",
		Price: 1.99,
		SKU: "xyz123",
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}