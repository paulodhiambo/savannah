package repositories

import (
	"backend/internal/models"
	"backend/pkg/database"
	"backend/pkg/logging"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository_Create(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewCustomerRepository(db, logger)

	customer := &models.Customer{Name: "John Doe", Code: "C123"}
	err := repo.Create(customer)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, customer.ID)
}

func TestCustomerRepository_Update(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewCustomerRepository(db, logger)

	customer := &models.Customer{Name: "John Doe", Code: "C123"}
	err := repo.Create(customer)
	assert.NoError(t, err)

	newName := "Jane Doe"
	customer.Name = newName
	err = repo.Update(customer)
	assert.NoError(t, err)
	updatedCustomer, err := repo.GetByID(customer.ID)
	assert.NoError(t, err)
	assert.Equal(t, newName, updatedCustomer.Name)
}

func TestCustomerRepository_GetAll(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewCustomerRepository(db, logger)

	customers := []models.Customer{
		{Name: "John Doe", Code: "C123"},
		{Name: "Jane Doe", Code: "C124"},
	}

	for _, c := range customers {
		err := repo.Create(&c)
		assert.NoError(t, err)
	}

	allCustomers, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, allCustomers, len(customers))
}
