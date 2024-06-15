package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewProduct(t *testing.T) {
	name := "Sample Product"
	description := "This is a sample product."
	price := 19.99
	product := NewProduct(name, description, price)

	assert.NotNil(t, product)
	assert.Equal(t, name, product.Name)
	assert.Equal(t, description, product.Description)
	assert.Equal(t, price, product.Price)
}

func TestProductFields(t *testing.T) {
	now := time.Now()

	product := &Product{
		Model: gorm.Model{
			ID:        1,
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:        "Test Product",
		Description: "Test Description",
		Price:       9.99,
	}

	assert.Equal(t, uint(1), product.ID)
	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, "Test Description", product.Description)
	assert.Equal(t, 9.99, product.Price)
	assert.WithinDuration(t, now, product.CreatedAt, time.Second)
	assert.WithinDuration(t, now, product.UpdatedAt, time.Second)
}
