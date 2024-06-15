package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCustomer(t *testing.T) {
	name := "John Doe"
	code := "C123"
	customer := NewCustomer(name, code)
	fmt.Printf("%s - %s", customer.CreatedAt, time.Now().Format(time.RFC3339))
	assert.NotNil(t, customer)
	assert.Equal(t, name, customer.Name)
	assert.Equal(t, code, customer.Code)
}

func TestCustomerFields(t *testing.T) {
	customer := &Customer{
		ID:   1,
		Name: "Jane Doe",
		Code: "C124",
	}

	assert.Equal(t, 1, customer.ID)
	assert.Equal(t, "Jane Doe", customer.Name)
	assert.Equal(t, "C124", customer.Code)
}
