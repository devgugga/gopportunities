package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Initialize() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080")
}
