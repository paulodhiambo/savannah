package repositories

import (
	"backend/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

type CustomerRepositoryImpl interface {
	Create(customer *models.Customer) error
	Update(customer *models.Customer) error
	GetByID(int) (*models.Customer, error)
	GetAll() ([]models.Customer, error)
}

func NewCustomerRepository(db *gorm.DB, logger *logrus.Logger) CustomerRepositoryImpl {
	return &CustomerRepository{DB: db, logger: logger}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	if err := r.DB.Create(customer).Error; err != nil {
		r.logger.Warnf("Error while creating customer: %v", err)
		return fmt.Errorf("failed to create customer: %v", err)
	}
	return nil
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
	if err := r.DB.Save(customer).Error; err != nil {
		r.logger.Warnf("Error while updating customer: %v", err)
		return fmt.Errorf("failed to update customer: %v", err)
	}
	return nil
}

func (r *CustomerRepository) GetByID(id int) (*models.Customer, error) {
	var customer models.Customer
	if err := r.DB.First(&customer, id).Error; err != nil {
		r.logger.Warnf("Error while getting customer: %v", err)
		return nil, fmt.Errorf("failed to get customer by ID: %v", err)
	}
	return &customer, nil
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
	var customers []models.Customer
	if err := r.DB.Find(&customers).Error; err != nil {
		r.logger.Warnf("Error while getting customers: %v", err)
		return nil, fmt.Errorf("failed to get all customers: %v", err)
	}
	return customers, nil
}
