package data

import "time"

type Product struct {
	ID int
	Name string
	Description string
	Price float32
	SKU string
	CreatedAt string
	UpdatedAt string
	DeletedAt string
}

type Products []*Product

func AllProducts() Products {
	return products
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