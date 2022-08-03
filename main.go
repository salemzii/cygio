package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/app"
)

func main() {

	router := gin.Default()

	router.GET("/", welcome)
	router.POST("/createalert", app.ReceiveAlert)

	router.Run()

}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}
