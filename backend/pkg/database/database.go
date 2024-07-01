package database

import (
	"backend/internal/config"
	"backend/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connect initializes the database connection based on the environment
func Connect(logger *logrus.Logger) error {

	var (
		err error
		dsn string
	)

	err = config.Load()
	if err != nil {
		return err
	}

	// Check the environment to determine the database driver and DSN
	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Africa/Nairobi", config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBUser, config.AppConfig.DBPassword, config.AppConfig.DBName)
	// Connect to the database
	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		logger.Warnf("failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto migrate the models
	if err := DB.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}); err != nil {
		logger.Warnf("failed to auto migrate models: %v", err)
		return fmt.Errorf("failed to auto migrate models: %v", err)
	}

	return nil
}

// Close closes the database connection
func Close(logger *logrus.Logger) error {
	db, err := DB.DB()
	if err != nil {
		logger.Warnf("failed to get database connection: %v", err)
		return fmt.Errorf("failed to get database connection: %v", err)
	}
	if db != nil {
		return db.Close()
	}
	return nil
}

// DropTables drops the database tables
func DropTables(logger *logrus.Logger) error {
	db := DB
	err := db.Migrator().DropTable(&models.Order{}, &models.Product{}, &models.Customer{})
	if err != nil {
		logger.Warnf("failed to drop tables: %v", err)
		return fmt.Errorf("failed to drop tables: %v", err)
	}
	return nil
}
