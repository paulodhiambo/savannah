package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// NewCustomer creates a new Customer instance
func NewCustomer(name, code string) *Customer {
	return &Customer{
		Name: name,
		Code: code,
	}
}
