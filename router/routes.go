package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// initializeRoutes provides all the application Routes.
func initializeRoutes(r *gin.Engine) {

	//Setting the api Path.
	v1 := r.Group("/api/v1/")
	{
		//This route get a job opening.
		v1.GET("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Show": "Opening",
			})
		})

		//This route get all job openings.
		v1.GET("/openings", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Show": "Openings",
			})
		})

		//This route create a job opening.
		v1.POST("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Create": "Opening",
			})
		})

		//This route update a job opening.
		v1.PUT("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Update": "Opening",
			})
		})

		//This route delete a job opening.
		v1.DELETE("/opening", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"Delete": "Opening",
			})
		})
	}
}
