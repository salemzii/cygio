package facebook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Messaging_product string `json:"messaging_product"`
	Recipient_type    string `json:"recipient_type"`
	To                string `json:"to"`
	Type              string `json:"type"`
	Text              *Text  `json:"text"`
	//Image             Image  `json:"image"`
	//Template          Template `json:"template"`
}

type Text struct {
	Preview_url bool   `json:"preview_url"`
	Body        string `json:"body"`
}

type Template struct {
	Name     string   `json:"name"`
	Language Language `json:"language"`
}

type Language struct {
	Code string `json:"code"`
}

type Image struct {
	Id string `json:"id"`
}

type MessageResponse struct {
	Messaging_product string    `json:"messaging_product"`
	Contacts          []Contact `json:"contacts"`
	Messages          []Msg     `json:"messages"`
}

type Contact struct {
	Input string `json:"input"`
	Wa_Id string `json:"wa_id"`
}

type Msg struct {
	Id string `json:"id"`
}

func SendMsg() {

	uri := "https://graph.facebook.com/v13.0/app_id/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = ""
	body.Type = "text"
	body.Text = &Text{Body: "hello salemzii from the goland Again", Preview_url: true}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{os.Getenv("META_ACCESS_TOKEN")},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}

/*
func SendMediaMsg() {

	uri := "https://graph.facebook.com/v13.0/app_id/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = ""
	body.Image = Image{}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{""},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}

/*
func SendTemplMsg() {

	uri := "https://graph.facebook.com/v13.0/app_id/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.To = ""
	body.Type = "template"
	body.Template = Template{Name: "hello_world", Language: Language{Code: "en_US"}}

	json_data, err := json.Marshal(body)

	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(json_data))

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{os.Getenv("META_ACCESS_TOKEN")},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// decode bytes response to a readable json

	var respJson MessageResponse
	err = json.Unmarshal(respByte, &respJson)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(respJson)
}
*/
