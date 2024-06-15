package repositories

import (
	"backend/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

type OrderRepositoryImpl interface {
	Create(order *models.Order) error
	Update(order *models.Order) error
	Delete(id int) error
	GetByID(id int) (*models.Order, error)
	GetAll() ([]models.Order, error)
	GetOrdersByUserID(userID int) ([]models.Order, error)
}

func NewOrderRepository(db *gorm.DB, logger *logrus.Logger) OrderRepositoryImpl {
	return &OrderRepository{DB: db, logger: logger}
}

func (r *OrderRepository) Create(order *models.Order) error {
	if err := r.DB.Create(order).Error; err != nil {
		r.logger.Warnf("failed to create order: %v", err)
		return fmt.Errorf("failed to create order: %v", err)
	}
	return nil
}

func (r *OrderRepository) Update(order *models.Order) error {
	if err := r.DB.Save(order).Error; err != nil {
		r.logger.Warnf("failed to update order: %v", err)
		return fmt.Errorf("failed to update order: %v", err)
	}
	return nil
}

func (r *OrderRepository) Delete(id int) error {
	if err := r.DB.Delete(&models.Order{}, id).Error; err != nil {
		r.logger.Warnf("failed to delete order: %v", err)
		return fmt.Errorf("failed to delete order: %v", err)
	}
	return nil
}

func (r *OrderRepository) GetByID(id int) (*models.Order, error) {
	var order models.Order
	if err := r.DB.First(&order, id).Error; err != nil {
		r.logger.Warnf("failed to get order: %v", err)
		return nil, fmt.Errorf("failed to get order by ID: %v", err)
	}
	return &order, nil
}

func (r *OrderRepository) GetAll() ([]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Find(&orders).Error; err != nil {
		r.logger.Warnf("failed to get all orders: %v", err)
		return nil, fmt.Errorf("failed to get all orders: %v", err)
	}
	return orders, nil
}

func (r *OrderRepository) GetOrdersByUserID(userID int) ([]models.Order, error) {
	var orders []models.Order
	if err := r.DB.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		r.logger.Warnf("failed to get orders by user ID: %v", err)
		return nil, fmt.Errorf("failed to get orders by user ID: %v", err)
	}
	return orders, nil
}
