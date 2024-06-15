package repositories

import (
	"backend/internal/models"
	"backend/pkg/database"
	"backend/pkg/logging"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductRepository_Create(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewProductRepository(db, logger)

	product := &models.Product{Name: "Test Product", Description: "Test Description", Price: 19.99}
	err := repo.Create(product)
	assert.NoError(t, err)

	assert.NotEqual(t, 0, product.ID)
}

func TestProductRepository_Update(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewProductRepository(db, logger)

	product := &models.Product{Name: "Test Product", Description: "Test Description", Price: 19.99}
	err := repo.Create(product)
	assert.NoError(t, err)

	newPrice := 29.99
	product.Price = newPrice
	err = repo.Update(product)
	assert.NoError(t, err)

	updatedProduct, err := repo.GetByID(int(product.ID))
	assert.NoError(t, err)
	assert.Equal(t, newPrice, updatedProduct.Price)
}

func TestProductRepository_Delete(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewProductRepository(db, logger)

	product := &models.Product{Name: "Test Product", Description: "Test Description", Price: 19.99}
	err := repo.Create(product)
	assert.NoError(t, err)

	err = repo.Delete(int(product.ID))
	assert.NoError(t, err)

	_, err = repo.GetByID(int(product.ID))
	assert.Error(t, err) // Product should not exist
}

func TestProductRepository_GetByID(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewProductRepository(db, logger)

	product := &models.Product{Name: "Test Product", Description: "Test Description", Price: 19.99}
	err := repo.Create(product)
	assert.NoError(t, err)

	fetchedProduct, err := repo.GetByID(int(product.ID))
	assert.NoError(t, err)
	assert.NotNil(t, fetchedProduct)
	assert.Equal(t, product.ID, fetchedProduct.ID)
}

func TestProductRepository_GetAll(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewProductRepository(db, logger)

	products := []*models.Product{
		{Name: "Product 1", Description: "Description 1", Price: 19.99},
		{Name: "Product 2", Description: "Description 2", Price: 29.99},
	}

	for _, p := range products {
		err := repo.Create(p)
		assert.NoError(t, err)
	}

	allProducts, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, allProducts, len(products))

	// Check if all product IDs match
	for i, p := range allProducts {
		assert.Equal(t, products[i].ID, p.ID)
	}
}
