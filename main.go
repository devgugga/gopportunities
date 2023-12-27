package main

import (
	"github.com/devgugga/gopportunities/config"
	_ "github.com/devgugga/gopportunities/docs"
	"github.com/devgugga/gopportunities/router"
)

var (
	logger *config.Logger
)

// @title Gopportunities API
// @version 1.0
// @description This documentation describes the GoPportunities API endpoints.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi
func main() {

	logger = config.GetLogger("main")

	// Initialize config
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing config: %v", err)
		return
	}

	// Initialize router
	router.Initialize()

}
