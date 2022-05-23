package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Message struct {
	Messaging_product string `json:"messaging_product"`
	Recipient_type    string `json:"recipient_type"`
	To                string `json:"to"`
	Type              string `json:"type"`
	Text              Text   `json:"text"`
	Image             Image  `json:"image"`
}

type Text struct {
	Preview_url bool   `json:"preview_url"`
	Body        string `json:"body"`
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

	uri := "https://graph.facebook.com/v13.0/101104952624494/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = "2347014327332"
	body.Type = "text"
	body.Text = Text{Body: "hello salemzii from the goland Again", Preview_url: true}

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
		"Authorization": []string{"Bearer EAAPawvj0HV0BAHeZAnzsRoCAZAy5BxwgzG3H1BX2bAKg8ranW02EMBeMoauWC6I9cbbpUrkDQ8HIo7aVyapljZBze90jaVx04RB6A6jkrvtQqZBIBSRIJkalUkv5aQxVDuYc9RXeaG3LQBEqKyZCTA0E9aEex3Wa3zoojVJe5ZACLj8JXk03mvxU2R0dfrxEmt4JdgE3ZAS3gZDZD"},
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

func SendMediaMsg() {

	uri := "https://graph.facebook.com/v13.0/101104952624494/messages"

	client := http.Client{}
	body := new(Message)

	body.Messaging_product = "whatsapp"
	body.Recipient_type = "individual"
	body.To = "2348053503763"
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
		"Authorization": []string{"Bearer EAAPawvj0HV0BAHeZAnzsRoCAZAy5BxwgzG3H1BX2bAKg8ranW02EMBeMoauWC6I9cbbpUrkDQ8HIo7aVyapljZBze90jaVx04RB6A6jkrvtQqZBIBSRIJkalUkv5aQxVDuYc9RXeaG3LQBEqKyZCTA0E9aEex3Wa3zoojVJe5ZACLj8JXk03mvxU2R0dfrxEmt4JdgE3ZAS3gZDZD"},
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
