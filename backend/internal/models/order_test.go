package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	productID := 1
	userId := 1
	quantity := 2
	total := 39.98
	order := NewOrder(productID, quantity, total, userId)

	assert.NotNil(t, order)
	assert.Equal(t, productID, order.ProductID)
	assert.Equal(t, userId, order.UserId)
	assert.Equal(t, quantity, order.Quantity)
	assert.Equal(t, total, order.Total)
}

func TestOrderFields(t *testing.T) {
	now := time.Now()

	order := &Order{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		ProductID: 1,
		UserId:    1,
		Quantity:  2,
		Total:     39.98,
	}

	assert.Equal(t, uint(1), order.ID)
	assert.Equal(t, 1, order.ProductID)
	assert.Equal(t, 1, order.UserId)
	assert.Equal(t, 2, order.Quantity)
	assert.Equal(t, 39.98, order.Total)
	assert.WithinDuration(t, now, order.CreatedAt, time.Second)
	assert.WithinDuration(t, now, order.UpdatedAt, time.Second)
}
