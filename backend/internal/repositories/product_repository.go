package repositories

import (
	"backend/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB     *gorm.DB
	logger *logrus.Logger
}

func NewProductRepository(db *gorm.DB, logger *logrus.Logger) *ProductRepository {
	return &ProductRepository{DB: db, logger: logger}
}

func (r *ProductRepository) Create(product *models.Product) error {
	if err := r.DB.Create(product).Error; err != nil {
		r.logger.Warnf("error creating product: %v", err)
		return fmt.Errorf("failed to create product: %v", err)
	}
	return nil
}

func (r *ProductRepository) Update(product *models.Product) error {
	if err := r.DB.Save(product).Error; err != nil {
		r.logger.Warnf("error updating product: %v", err)
		return fmt.Errorf("failed to update product: %v", err)
	}
	return nil
}

func (r *ProductRepository) Delete(id int) error {
	if err := r.DB.Delete(&models.Product{}, id).Error; err != nil {
		r.logger.Warnf("error deleting product: %v", err)
		return fmt.Errorf("failed to delete product: %v", err)
	}
	return nil
}

func (r *ProductRepository) GetByID(id int) (*models.Product, error) {
	var product models.Product
	if err := r.DB.First(&product, id).Error; err != nil {
		r.logger.Warnf("error getting product: %v", err)
		return nil, fmt.Errorf("failed to get product by ID: %v", err)
	}
	return &product, nil
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	if err := r.DB.Find(&products).Error; err != nil {
		r.logger.Warnf("error getting products: %v", err)
		return nil, fmt.Errorf("failed to get all products: %v", err)
	}
	return products, nil
}
