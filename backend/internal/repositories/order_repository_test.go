package repositories

import (
	"backend/internal/models"
	"backend/pkg/database"
	"backend/pkg/logging"
	_ "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderRepository_Create(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	order := &models.Order{ProductID: 1, Quantity: 2, Total: 39.98, UserId: 1}
	err := repo.Create(order)
	assert.NoError(t, err)

	assert.NotEqual(t, 0, order.ID)
}

func TestOrderRepository_Update(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	order := &models.Order{ProductID: 1, Quantity: 2, Total: 39.98, UserId: 1}
	err := repo.Create(order)
	assert.NoError(t, err)

	newQuantity := 3
	order.Quantity = newQuantity
	err = repo.Update(order)
	assert.NoError(t, err)

	updatedOrder, err := repo.GetByID(int(order.ID))
	assert.NoError(t, err)
	assert.Equal(t, newQuantity, updatedOrder.Quantity)
}

func TestOrderRepository_Delete(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	order := &models.Order{ProductID: 1, Quantity: 2, Total: 39.98, UserId: 1}
	err := repo.Create(order)
	assert.NoError(t, err)

	err = repo.Delete(int(order.ID))
	assert.NoError(t, err)

	_, err = repo.GetByID(int(order.ID))
	assert.Error(t, err) // Order should not exist
}

func TestOrderRepository_GetByID(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	order := &models.Order{ProductID: 1, Quantity: 2, Total: 39.98, UserId: 1}
	err := repo.Create(order)
	assert.NoError(t, err)

	fetchedOrder, err := repo.GetByID(int(order.ID))
	assert.NoError(t, err)
	assert.NotNil(t, fetchedOrder)
	assert.Equal(t, order.ID, fetchedOrder.ID)
}

func TestOrderRepository_GetAll(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	orders := []*models.Order{
		{ProductID: 1, Quantity: 2, Total: 39.98, UserId: 1},
		{ProductID: 2, Quantity: 3, Total: 59.97, UserId: 2},
	}

	for _, o := range orders {
		err := repo.Create(o)
		assert.NoError(t, err)
	}

	allOrders, err := repo.GetAll()
	assert.NoError(t, err)
	assert.Len(t, allOrders, len(orders))

	// Check if all order IDs match
	for i, o := range allOrders {
		assert.Equal(t, orders[i].ID, o.ID)
	}
}

func TestOrderRepository_GetOrdersByUserID(t *testing.T) {
	logger := logging.GetLogger()
	defer func() {
		_ = database.DropTables(logger)
		_ = database.Close(logger)
	}()

	_ = database.Connect(logger)
	db := database.DB

	repo := NewOrderRepository(db, logger)

	// Create test data
	orders := []models.Order{
		{ProductID: 1, Quantity: 2, UserId: 1, Total: 20.0},
		{ProductID: 2, Quantity: 1, UserId: 1, Total: 10.0},
		{ProductID: 3, Quantity: 5, UserId: 2, Total: 50.0},
	}

	for _, order := range orders {
		err := repo.Create(&order)
		assert.NoError(t, err)
	}
}
