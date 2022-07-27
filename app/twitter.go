package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	URL = "https://api.twitter.com/2"
)

func GetTweetById(id int) *GetTweetResponse {
	endpoint := fmt.Sprintf("/tweets?ids=%d", id)

	httpClient := cred.GetClientToken()

	path := URL + endpoint
	resp, err := httpClient.Get(path)
	if err != nil {
		log.Println("Error: ", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error: ", err)
	}
	var results GetTweetResponse
	if err := json.Unmarshal(body, &results); err != nil {
		log.Println(err)
	}

	return &results
}

func CreateTweet(text string) (response *CreateTweetResponse, err error) {
	endpoint := "/tweets"
	path := URL + endpoint
	tweet := Createtweetrequest{
		Text: text,
	}
	data, err := json.Marshal(&tweet)
	if err != nil {
		log.Println("Error encoding to json", err)
	}

	httpClient := cred.GetClientToken()
	res, err := httpClient.Post(path, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error making request: ", err)
	}

	fmt.Println(res.Status)
	// decode the response body to a  bytes
	respByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var resp CreateTweetResponse

	if err := json.Unmarshal(respByte, &resp); err != nil {
		log.Println("Error decoding response::: ", err)
	}

	return &resp, nil
}
