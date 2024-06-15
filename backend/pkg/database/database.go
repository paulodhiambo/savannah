package database

import (
	"backend/internal/models"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

// Connect initializes the database connection based on the environment
func Connect(logger *logrus.Logger) error {
	var (
		err error
		dsn string
	)

	// Check the environment to determine the database driver and DSN
	//env := os.Getenv("ENVIRONMENT")
	env := "development"
	if env == "production" {
		dbHost := os.Getenv("DB_HOST")
		dbPort := os.Getenv("DB_PORT")
		dbUser := os.Getenv("DB_USER")
		dbPassword := os.Getenv("DB_PASSWORD")
		dbName := os.Getenv("DB_NAME")

		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
		// Connect to the database
		if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			logger.Warnf("failed to connect to database: %v", err)
			return fmt.Errorf("failed to connect to database: %v", err)
		}
	} else {
		dsn = "test.db" // SQLite database file for testing
		if DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{}); err != nil {
			logger.Warnf("failed to connect to database: %v", err)
			return fmt.Errorf("failed to connect to database: %v", err)
		}
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
