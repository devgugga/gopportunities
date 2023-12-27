package config

import (
	"github.com/devgugga/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Initialize postgres
	db, err := gorm.Open(postgres.Open("host=localhost user=mustty password=mypass132 dbname=gopportunities port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing postgres: %v", err)
		return nil, err
	}

	// Check if the table already exists
	if db.Migrator().HasTable(&schemas.Opening{}) {
		logger.Info("Database already exists proceeding...")
		return db, nil
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Error migrating postgres: %v", err)
		return nil, err
	}

	return db, nil
}
