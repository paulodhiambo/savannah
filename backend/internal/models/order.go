package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	UserId    int     `json:"user_id"`
	Total     float64 `json:"total"`
}

// NewOrder creates a new Order instance
func NewOrder(productID, quantity int, total float64, userId int) *Order {
	return &Order{
		ProductID: productID,
		Quantity:  quantity,
		Total:     total,
		UserId:    userId,
	}
}
