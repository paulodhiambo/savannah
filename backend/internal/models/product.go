package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// NewProduct creates a new Product instance
func NewProduct(name, description string, price float64) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Price:       price,
	}
}
