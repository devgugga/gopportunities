package router

import (
	"github.com/devgugga/gopportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// initializeRoutes provides all the application Routes.
func initializeRoutes(r *gin.Engine) {

	//Set the api Path.
	v1 := r.Group("/api/v1/")
	{
		//This route get a job opening.
		v1.GET("/opening", handler.ShowOpeningHandler)

		//This route get all job openings.
		v1.GET("/openings", handler.ShowAllOpeningsHandler)

		//This route create a job opening.
		v1.POST("/opening", handler.CreateOpeningHandler)

		//This route update a job opening.
		v1.PUT("/opening", handler.UpdateOpeningHandler)

		//This route delete a job opening.
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
