package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/app"
)

func main() {

	alert := app.Alert{
		Body:      "Alert Body with context",
		Platforms: []app.Platform{{Name: "twitter"}, {Name: "facebook"}},
		Urls: []app.Url{{Uri: "http://127.0.0.1:8000/webhook1"}, {Uri: "http://127.0.0.1:8000/webhook2"},
			{Uri: "http://127.0.0.1:8000/webhook3"}},
	}
	app.CreateAlerts(alert)

	/*
		router := gin.Default()

		router.GET("/", welcome)
		router.POST("/createalert", app.ReceiveAlert)

		router.Run()
	*/

}

func welcome(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "Hello welcome to Franka webhook",
	})
}
