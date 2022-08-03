package app

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/salemzii/cygio/facebook"
	"github.com/salemzii/cygio/twitter"
)

var (
	Wg     sync.WaitGroup
	Endpwg sync.WaitGroup
)

func ReceiveAlert(c *gin.Context) {

	var alert Alert
	if c.Request.Body != nil {
		c.BindJSON(&alert)
		log.Println(alert)
	}
	log.Println("Dispensing alerts to all listed platforms and urls")
	defer CreateAlerts(alert)

	c.JSON(200, gin.H{
		"success": "Alerts distributed successfully",
	})
}

func CreateAlerts(alert Alert) {

	Platforms := alert.Platforms
	Urls := alert.Urls
	text := alert.Body

	Wg.Add(len(Platforms))
	for _, v := range Platforms {
		switch v.Name {

		case "twitter":
			log.Println("Creating alert on twitter")
			go func(text string) {
				defer Wg.Done()
				twitter.CreateTweet(text)
			}(text)

		case "facebook":
			log.Println("Creating alert on facebook")
			go func(text string) {
				defer Wg.Done()
				facebook.PagePost(text)
			}(text)

		}
	}

	Wg.Wait()

	SendToEndpoints(Urls, text)

}

type AlertBody struct {
	Body string `json:"body"`
}

func SendToEndpoints(Urls []Url, content string) {
	alertBody := AlertBody{Body: content}
	data, err := json.Marshal(&alertBody)
	if err != nil {
		log.Println("Error encoding to json", err)
	}

	Endpwg.Add(len(Urls))

	for _, v := range Urls {
		go func(url Url) {
			defer Endpwg.Done()
			resp, err := http.Post(url.Uri, "application/json", bytes.NewBuffer(data))
			if err != nil {
				log.Printf("Error posting data to %s", err)
			}
			defer resp.Body.Close()
			log.Println(resp.Status)
		}(v)
	}

	Endpwg.Wait()
}
