package app

import (
	"github.com/gin-gonic/gin"
	"github.com/juankis/registry/api/src/controllers"
)

//Start function to start up proyect
func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/registry", controllers.InsertRegistry)
	r.POST("/registry", controllers.GetRegistry)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
