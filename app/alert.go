package app

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Alert struct {
	Headline  string     `json:"headline"`
	Media     []Media    `json:"media"`
	Body      string     `json:"body"`
	Platforms []Platform `json:"platforms"`
}

type Media struct {
	Description string `json:"description"`
	Url         string `json:"url"`
}

type Platform struct {
	Name string `json:"name"`
}

func ReceiveAlert(c *gin.Context) {

	var alert Alert
	if c.Request.Body != nil {
		c.BindJSON(&alert)
		log.Print(alert)
	}
}
