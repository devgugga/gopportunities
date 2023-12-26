package router

import (
	"github.com/gin-gonic/gin"
)

// Initialize the API server on port 8080
func Initialize() {
	//Initialize Router
	r := gin.Default()

	//Initialize Routes
	initializeRoutes(r)

	//Run the server
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
