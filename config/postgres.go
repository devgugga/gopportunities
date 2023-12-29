package config

import (
	"fmt"
	"os/exec"

	"github.com/devgugga/gopportunities/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func createDockerContainer() {
	fmt.Println("Creating docker container for postgres...")
	cmd := exec.Command("docker-compose", "up", "-d")
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error creating docker container: %v", err)
		return
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Docker container created successfully, please start the server again")
	fmt.Println("-----------------------------------")
}

func InitPostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Initialize postgres
	db, err := gorm.Open(postgres.Open("host=localhost user=mustty password=mypass132 dbname=gopportunities port=5432"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error initializing postgres: %v", err)
		createDockerContainer()
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
