package facebook

import (
	"errors"
	"sync"
)

var (
	UnmarshalError = errors.New("Unmarshal Error")
	Wg             sync.WaitGroup
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
